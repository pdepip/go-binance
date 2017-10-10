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


