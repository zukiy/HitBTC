package hitbtc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Request(t *testing.T) {
	var testCases = []struct {
		method, url        string
		responseStatusCode int
		responseBody       []byte
		expectResult       []byte
		expectError        error
	}{
		{
			http.MethodGet,
			cAPIURLRest + cFetchSymbols,
			http.StatusOK,
			nil,
			[]byte{},
			nil,
		},
		{
			http.MethodGet,
			"",
			http.StatusOK,
			nil,
			[]byte{},
			nil,
		},
	}

	for _, tc := range testCases {
		httpClientMock := new(httpMock)
		c := NewWithHTTPClient(httpClientMock)

		request, _ := http.NewRequest(tc.method, tc.url, nil)

		httpClientMock.On(`Do`, request).Return(
			makeMockResponse(tc.responseStatusCode, tc.responseBody),
			nil,
		)

		result, err := c.request(tc.method, tc.url, nil)
		assert.Equal(t, tc.expectResult, result)
		assert.Equal(t, tc.expectError, err)
	}

}

func TestClient_FetchSymbol(t *testing.T) {
	var testCases = []struct {
		symbol             string
		responseStatusCode int
		responseBody       []byte
		expectError        error
		expectResult       *Symbol
	}{
		{
			symbol:             "MRSUSD",
			responseStatusCode: 200,
			responseBody:       []byte(`{"id":"MRSUSD","baseCurrency":"MRS","quoteCurrency":"USD","quantityIncrement":"100","tickSize":"0.000001","takeLiquidityRate":"0.001","provideLiquidityRate":"-0.0001","feeCurrency":"USD"}`),
			expectError:        nil,
			expectResult: &Symbol{
				ID:                   "MRSUSD",
				BaseCurrency:         "MRS",
				QuoteCurrency:        "USD",
				QuantityIncrement:    100.,
				TickSize:             0.000001,
				TakeLiquidityRate:    0.001,
				ProvideLiquidityRate: -0.0001,
				FeeCurrency:          "USD",
			},
		},
		{
			symbol:             "BAD_SYMBOL_VALUE",
			responseStatusCode: 400,
			responseBody:       []byte(`{"error": {"code": 2001, "message": "Symbol not found", "description": "Try get /api/2/public/symbol, to get list of all available symbols."}}`),
			expectError:        errors.New("[2001] Symbol not found"),
			expectResult:       nil,
		},
	}

	httpClientMock := new(httpMock)
	c := NewWithHTTPClient(httpClientMock)

	for _, tc := range testCases {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(`https://api.hitbtc.com/api/2/public/symbol/%s`, tc.symbol), nil)

		httpClientMock.On(`Do`, request).Return(
			makeMockResponse(tc.responseStatusCode, tc.responseBody),
			nil,
		)

		result, err := c.FetchSymbol(tc.symbol)

		assert.Equal(t, tc.expectResult, result)
		assert.Equal(t, tc.expectError, err)
	}
}

func makeMockResponse(code int, body []byte) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
	}
}

func TestClient_FetchSymbols(t *testing.T) {
	var testCases = []struct {
		responseStatusCode int
		responseBody       []byte
		expectError        error
		expectResult       Symbols
		expectResultCount  int
	}{
		{
			responseStatusCode: 200,
			responseBody:       []byte(`[{"id":"PAXETH","baseCurrency":"PAX","quoteCurrency":"ETH","quantityIncrement":"1","tickSize":"0.00001","takeLiquidityRate":"0.001","provideLiquidityRate":"-0.0001","feeCurrency":"ETH"},{"id":"PAXUSD","baseCurrency":"PAX","quoteCurrency":"USD","quantityIncrement":"1","tickSize":"0.0001","takeLiquidityRate":"0.001","provideLiquidityRate":"-0.0001","feeCurrency":"USD"},{"id":"PAXEOS","baseCurrency":"PAX","quoteCurrency":"EOS","quantityIncrement":"1","tickSize":"0.00001","takeLiquidityRate":"0.001","provideLiquidityRate":"-0.0001","feeCurrency":"EOS"}]`),
			expectError:        nil,
			expectResult: Symbols{
				{ID: "PAXETH", BaseCurrency: "PAX", QuoteCurrency: "ETH", QuantityIncrement: 1, TickSize: 0.00001, TakeLiquidityRate: 0.001, ProvideLiquidityRate: -0.0001, FeeCurrency: "ETH"},
				{ID: "PAXUSD", BaseCurrency: "PAX", QuoteCurrency: "USD", QuantityIncrement: 1, TickSize: 0.0001, TakeLiquidityRate: 0.001, ProvideLiquidityRate: -0.0001, FeeCurrency: "USD"},
				{ID: "PAXEOS", BaseCurrency: "PAX", QuoteCurrency: "EOS", QuantityIncrement: 1, TickSize: 0.00001, TakeLiquidityRate: 0.001, ProvideLiquidityRate: -0.0001, FeeCurrency: "EOS"},
			},
			expectResultCount: 3,
		},
		{
			responseStatusCode: 504,
			responseBody:       []byte(`{"error": {"code": 504, "message": "Gateway Timeout", "description": "Check the result of your request later"}}`),
			expectError:        errors.New("[504] Gateway Timeout"),
			expectResult:       make(Symbols, 0),
			expectResultCount:  0,
		},
	}

	for _, tc := range testCases {
		httpClientMock := new(httpMock)
		c := NewWithHTTPClient(httpClientMock)

		request, _ := http.NewRequest(http.MethodGet, "https://api.hitbtc.com/api/2/public/symbol", nil)

		httpClientMock.On(`Do`, request).Return(
			makeMockResponse(tc.responseStatusCode, tc.responseBody),
			nil,
		)

		result, err := c.FetchSymbols()

		assert.Equal(t, tc.expectError, err)
		assert.Equal(t, tc.expectResult, result)
		assert.Equal(t, tc.expectResultCount, len(result))
	}
}

func TestClient_SubscribeToOrderBookFor(t *testing.T) {
	c := New()
	books, done, err := c.SubscribeToOrderBookFor("SBTCUSDT")
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second * 5)
		done <- struct{}{}
	}()

	for {
		select {
		case b := <-books:
			fmt.Printf("%+v\n", b)
		}
	}
}
