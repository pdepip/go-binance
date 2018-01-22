/*

   Stores request structs & their respectivate validation functions for Binance API

*/

package binance

import (
	"errors"
)

// Input for: POST /api/v3/order
type LimitOrder struct {
	Symbol      string
	Side        string
	Type        string
	TimeInForce string
	Quantity    float64
	Price       float64
	RecvWindow  int64
}

// Validating a Limit Order
func (l *LimitOrder) ValidateLimitOrder() error {
	switch {
	case len(l.Symbol) == 0:
		return errors.New("Order must contain a symbol")
	case !OrderSideEnum[l.Side]:
		return errors.New("Invalid or empty order side")
	case l.Type != "LIMIT":
		return errors.New("Invalid LIMIT order type")
	case !OrderTIFEnum[l.TimeInForce]:
		return errors.New("Invalid or empty order timeInForce")
	case l.Quantity <= 0.0:
		return errors.New("Invalid or empty order quantity")
	case l.Price <= 0.0:
		return errors.New("Invalid or empty order price")
	case l.RecvWindow == 0:
		l.RecvWindow = 5000
		return nil
	default:
		return nil
	}
}

type MarketOrder struct {
	Symbol     string
	Side       string
	Type       string
	Quantity   float64
	RecvWindow int64
}

func (m *MarketOrder) ValidateMarketOrder() error {
	switch {
	case len(m.Symbol) == 0:
		return errors.New("Order must contain a symbol")
	case !OrderSideEnum[m.Side]:
		return errors.New("Invalid or empty order side")
	case m.Quantity <= 0.0:
		return errors.New("Invalid or empty order quantity")
	case m.RecvWindow == 0:
		m.RecvWindow = 5000
		return nil
	default:
		return nil
	}
}

// Input for: GET & DELETE /api/v3/order
type OrderQuery struct {
	Symbol     string
	OrderId    int64
	RecvWindow int64
}

func (q *OrderQuery) ValidateOrderQuery() error {
	switch {
	case len(q.Symbol) == 0:
		return errors.New("OrderQuery must contain a symbol")
	case q.OrderId == 0:
		return errors.New("OrderQuery must contain an OrderId")
	case q.RecvWindow == 0:
		q.RecvWindow = 5000
		return nil
	default:
		return nil
	}
}

// Input for: GET /api/v3/openOrders
type OpenOrdersQuery struct {
	Symbol     string
	RecvWindow int64
}

func (q *OpenOrdersQuery) ValidateOpenOrdersQuery() error {
	switch {
	case len(q.Symbol) == 0:
		return errors.New("OpenOrderQuery must contain a symbol")
	case q.RecvWindow == 0:
		q.RecvWindow = 5000
		return nil
	default:
		return nil
	}
}

// Input for: GET /api/v3/allOrders
type AllOrdersQuery struct {
	Symbol     string
	OrderId    int64
	Limit      int64
	RecvWindow int64
}

func (q *AllOrdersQuery) ValidateAllOrdersQuery() error {
	if len(q.Symbol) == 0 {
		return errors.New("AllOrderQuery must contain a symbol")
	}

	if q.Limit > 500 {
		return errors.New("AllOrderQuery limit max is 500")
	}

	if q.RecvWindow == 0 {
		q.RecvWindow = 5000
	}

	if q.Limit == 0 {
		q.Limit = 500
	}

	return nil
}
