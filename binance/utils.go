/*
   Util Functions for Binance Api Wrapper
*/
package binance

import (
	"time"
)

func unixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func recvWindow(d time.Duration) int64 {
	return int64(d) / int64(time.Millisecond)
}
