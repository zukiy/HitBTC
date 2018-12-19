package hitbtc

import "fmt"

// Err struct
// easyjson:json
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

// Symbol model
// easyjson:json
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

// SubscribeOrderBookResponse ws response
//easyjson:json
type SubscribeOrderBookResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Ask       []Order `json:"ask"`
		Bid       []Order `json:"bid"`
		Symbol    string  `json:"symbol"`
		Sequence  int64   `json:"sequence"`
		Timestamp string  `json:"timestamp"`
	} `json:"params"`
}

// SubscribeOrderBookResponseChan channel
type SubscribeOrderBookResponseChan chan SubscribeOrderBookResponse

// Order ...
//easyjson:json
type Order struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size,string"`
}

//easyjson:json
type subscribeResult struct {
	JSONRPC string `json:"jsonrpc"`
	Result  bool   `json:"result"`
}

// DoneChan break listener
type DoneChan chan struct{}

// ErrorHandler handles errors
type ErrorHandler func(err error)
