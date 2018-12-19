package hitbtc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type httpFetcher interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client engine
type Client struct {
	httpClient httpFetcher
}

// New create and return Client instance
func New() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// NewWithHTTPClient create and return Client instance with http client
func NewWithHTTPClient(httpClient httpFetcher) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// FetchSymbols return symbols list
func (c *Client) FetchSymbols() (Symbols, error) {
	list := make(Symbols, 0)

	bytes, err := c.request(http.MethodGet, cApiUrlRest+cFetchSymbols, nil)
	if err != nil {
		return list, err
	}

	err = list.UnmarshalJSON(bytes)
	if err != nil {
		return nil, err
	}

	return list, err
}

// FetchSymbol return symbol
func (c *Client) FetchSymbol(symbol string) (*Symbol, error) {
	bytes, err := c.request(http.MethodGet, cApiUrlRest+fmt.Sprintf(cFetchSymbol, symbol), nil)
	if err != nil {
		return nil, err
	}

	s := &Symbol{}
	err = s.UnmarshalJSON(bytes)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// FetchSymbol return symbol
func (c *Client) SubscribeToOrderBookFor(symbol string) (oBooks SubscribeOrderBookResponseChan, doneChan DoneChan, err error) {
	u := url.URL{
		Scheme: cApiStreamingScheme,
		Host:   cApiUrlStreamingHost,
		Path:   cApiUrlStreamingPath,
	}

	// init websocket
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return
	}

	// prepare payload
	requestParams, err := requestParams{
		Method: CSubscribeOrderbook,
		Params: params{
			Symbol: symbol,
		},
	}.MarshalJSON()

	if err != nil {
		return
	}

	// subscribe
	err = conn.WriteMessage(websocket.TextMessage, requestParams)
	if err != nil {
		return
	}

	// check for subscribe success
	_, r, err := conn.ReadMessage()
	if err != nil {
		return
	}

	subscribeResult := subscribeResult{}
	err = subscribeResult.UnmarshalJSON(r)
	if err != nil {
		return
	}

	if !subscribeResult.Result {
		err = fmt.Errorf("cant subscribe to orderbook for %s", symbol)
		return
	}

	doneChan = make(DoneChan)
	oBooks = make(SubscribeOrderBookResponseChan)

	go func() {
		defer conn.Close()
		for {
			select {
			case <-doneChan:
				return
			default:
				_, r, err := conn.ReadMessage()
				if err != nil {
					continue
				}

				response := SubscribeOrderBookResponse{}
				err = response.UnmarshalJSON(r)
				if err != nil {
					continue
				}

				select {
				case oBooks <- response:
				default:
				}
			}
		}
	}()

	return
}

func (c *Client) request(method, url string, body []byte) (response []byte, err error) {
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return response, err
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusOK {
		e := Err{}
		err = e.UnmarshalJSON(response)
		if err != nil {
			return response, err
		}
		return response, errors.New(e.Error())
	}

	return response, err
}
