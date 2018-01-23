/*

   Stores request strcuts & validation functions for API functions in account.go

*/

package binance

import (
	"errors"
)

// Input for: GET /api/v1/depth
type OrderBookQuery struct {
	Symbol string
	Limit  int64
}

func (q *OrderBookQuery) ValidateOrderBookQuery() error {
	switch {
	case len(q.Symbol) == 0:
		return errors.New("OrderBookQuery must contain a symbol")
	case q.Limit == 0:
		q.Limit = 100
		return nil
	default:
		return nil
	}
}

// Input for: GET /api/v1/24hr & /api/v1/aggTrades
type SymbolQuery struct {
	Symbol string
}

func (q *SymbolQuery) ValidateSymbolQuery() error {
	if len(q.Symbol) == 0 {
		return errors.New("Invalid or Empty Symbol")
	} else {
		return nil
	}
}

// Input for: Get /api/v1/klines
type KlineQuery struct {
	Symbol   string
	Interval string
	Limit    int64
}

func (q *KlineQuery) ValidateKlineQuery() error {
	switch {
	case len(q.Symbol) == 0:
		return errors.New("KlineQuery requires a symbol")
	case !IntervalEnum[q.Interval]:
		return errors.New("Invalid Kline Interval")
	case q.Limit == 0:
		q.Limit = 500
		return nil
	default:
		return nil
	}
}
