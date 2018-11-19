package hitbtc

import "fmt"

// easyjson:json
// Err struct
type Err struct {
	Data struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"error"`
}

func (e Err) Error() string {
	return fmt.Sprintf("[%d] %s", e.Data.Code, e.Data.Message)
}

// easyjson:json
// Symbol model
type Symbol struct {
	ID                   string  `json:"id"`
	BaseCurrency         string  `json:"baseCurrency"`
	QuoteCurrency        string  `json:"quoteCurrency"`
	QuantityIncrement    float64 `json:"quantityIncrement,string"`
	TickSize             float64 `json:"tickSize,string"`
	TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
	ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
	FeeCurrency          string  `json:"feeCurrency"`
}

// Symbols list
//easyjson:json
type Symbols []Symbol

type Ask struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size,string"`
}

//easyjson:json
type params struct {
	Symbol string `json:"symbol"`
}

//easyjson:json
type requestParams struct {
	Method string `json:"method"`
	Params params `json:"params"`
	ID     int    `json:"id"`
}

//easyjson:json
type SubscribeOrderBookResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Ask       []OBook `json:"ask"`
		Bid       []OBook `json:"bid"`
		Symbol    string  `json:"symbol"`
		Sequence  int64   `json:"sequence"`
		Timestamp string  `json:"timestamp"`
	} `json:"params"`
}

// SubscribeOrderBookResponseChan channel
type SubscribeOrderBookResponseChan chan SubscribeOrderBookResponse

//easyjson:json
type OBook struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size,string"`
}

//easyjson:json
type subscribeResult struct {
	JsonRPC string `json:"jsonrpc"`
	Result  bool   `json:"result"`
}

// DoneChan break listener
type DoneChan chan struct{}
