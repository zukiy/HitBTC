package HitBTC

const (
	// http
	cApiUrlRest   = `https://api.hitbtc.com/api/2`
	cFetchSymbols = `/public/symbol`
	cFetchSymbol  = `/public/symbol/%s`

	// ws
	cApiStreamingScheme  = `wss`
	cApiUrlStreamingHost = `api.hitbtc.com`
	cApiUrlStreamingPath = `api/2/ws`

	cSubscribeOrderbook = "subscribeOrderbook"
)
