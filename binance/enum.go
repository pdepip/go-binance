/*

   Stores ENUM values for parameters in API requests

*/

package binance

var OrderSideEnum = map[string]bool{
	"BUY":  true,
	"SELL": true,
}

var OrderTypeEnum = map[string]bool{
	"LIMIT":  true,
	"MARKET": true,
}

var OrderTIFEnum = map[string]bool{
	"GTC": true,
	"IOC": true,
}

var IntervalEnum = map[string]bool{
	"1m":  true,
	"3m":  true,
	"5m":  true,
	"15m": true,
	"30m": true,
	"1h":  true,
	"2h":  true,
	"4h":  true,
	"6h":  true,
	"8h":  true,
	"12h": true,
	"1d":  true,
	"3d":  true,
	"1w":  true,
	"1M":  true,
}
