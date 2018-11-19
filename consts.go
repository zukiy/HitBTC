package hitbtc

const (
	// http
	cApiUrlRest   = `https://api.hitbtc.com/api/2`
	cFetchSymbols = `/public/symbol`
	cFetchSymbol  = `/public/symbol/%s`

	// ws
	cApiStreamingScheme  = `wss`
	cApiUrlStreamingHost = `api.hitbtc.com`
	cApiUrlStreamingPath = `api/2/ws`

	CSubscribeOrderbook = "subscribeOrderbook"

	// CSnapshotOrderbook notification snapshot
	CSnapshotOrderbookMethod = "snapshotOrderbook"

	// CSnapshotOrderbook notification update
	CUpdateOrderbookMethod = "updateOrderbook"
)
