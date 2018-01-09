/*

   Stores request strcuts & validation functions for API functions in account.go

*/

package binance

import (
	"errors"
	"time"
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

//SymbolQuery Input for: GET /api/v1/24hr
type SymbolQuery struct {
	Symbol string
}

//ValidateSymbolQuery validate query input
func (q *SymbolQuery) ValidateSymbolQuery() error {
	if len(q.Symbol) == 0 {
		return errors.New("Invalid or Empty Symbol")
	}
	return nil
}

//AggTradesQuery Input for: GET /api/v1/aggTrades
type AggTradesQuery struct {
	Symbol    string     `url:"symbol"`
	FromID    int64      `url:"fromId,omitempty"`
	StartTime *time.Time `url:"startTime,omitempty"`
	EndTime   *time.Time `url:"endTime,omitempty"`
	Limit     int        `url:"limit,omitempty"`
}

//ValidateSymbolQuery validate query input
func (q *AggTradesQuery) ValidateSymbolQuery() error {
	if len(q.Symbol) == 0 {
		return errors.New("Invalid or Empty Symbol")
	}
	return nil
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
