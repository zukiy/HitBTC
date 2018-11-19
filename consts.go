package hitbtc

const (
	// http
	cAPIURLRest   = `https://api.hitbtc.com/api/2`
	cFetchSymbols = `/public/symbol`
	cFetchSymbol  = `/public/symbol/%s`

	// ws
	cAPIStreamingScheme  = `wss`
	cAPIURLStreamingHost = `api.hitbtc.com`
	cAPIUrlStreamingPath = `api/2/ws`

	// CSubscribeOrderbook method name
	CSubscribeOrderbook = "subscribeOrderbook"

	// CSnapshotOrderbookMethod notification snapshot
	CSnapshotOrderbookMethod = "snapshotOrderbook"

	// CUpdateOrderbookMethod notification update
	CUpdateOrderbookMethod = "updateOrderbook"
)
