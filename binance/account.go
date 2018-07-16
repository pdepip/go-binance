/*

   account.go
       Account (Signed) Endpoints for Binance Exchange API

*/
package binance

import (
	"fmt"
)

// Get Basic Account Information
func (b *Binance) GetAccountInfo() (account Account, err error) {

	reqUrl := fmt.Sprintf("api/v3/account")

	_, err = b.client.do("GET", reqUrl, "", true, &account)
	if err != nil {
		return
	}

	return
}

// Filter Basic Account Information To Retrieve Current Holdings
func (b *Binance) GetPositions() (positions []Balance, err error) {

	reqUrl := fmt.Sprintf("api/v3/account")
	account := Account{}

	_, err = b.client.do("GET", reqUrl, "", true, &account)
	if err != nil {
		return
	}

	positions = make([]Balance, len(account.Balances))
	i := 0

	for _, balance := range account.Balances {
		if balance.Free != 0.0 || balance.Locked != 0.0 {
			positions[i] = balance
			i++
		}
	}

	return positions[:i], nil
}

// Place a Limit Order
func (b *Binance) PlaceLimitOrder(l LimitOrder) (res PlacedOrder, err error) {

	err = l.ValidateLimitOrder()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v3/order?symbol=%s&side=%s&type=%s&timeInForce=%s&quantity=%f&price=%.8f&recvWindow=%d", l.Symbol, l.Side, l.Type, l.TimeInForce, l.Quantity, l.Price, l.RecvWindow)

	_, err = b.client.do("POST", reqUrl, "", true, &res)
	if err != nil {
		return
	}

	return
}

// Place a Market Order
func (b *Binance) PlaceMarketOrder(m MarketOrder) (res PlacedOrder, err error) {

	err = m.ValidateMarketOrder()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v3/order?symbol=%s&side=%s&type=%s&quantity=%f&recvWindow=%d", m.Symbol, m.Side, m.Type, m.Quantity, m.RecvWindow)

	_, err = b.client.do("POST", reqUrl, "", true, &res)
	if err != nil {
		return
	}

	return
}

// Cancel an Order
func (b *Binance) CancelOrder(query OrderQuery) (order CanceledOrder, err error) {

	err = query.ValidateOrderQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v3/order?symbol=%s&orderId=%d&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

	_, err = b.client.do("DELETE", reqUrl, "", true, &order)
	if err != nil {
		return
	}

	return
}

// Check the Status of an Order
func (b *Binance) CheckOrder(query OrderQuery) (status OrderStatus, err error) {

	err = query.ValidateOrderQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("api/v3/order?symbol=%s&orderId=%d&recvWindow=%d", query.Symbol, query.OrderId, query.RecvWindow)

	_, err = b.client.do("GET", reqUrl, "", true, &status)
	if err != nil {
		return
	}

	return
}

// Retrieve All Open Orders
func (b *Binance) GetAllOpenOrders() (orders []OrderStatus, err error) {
	_, err = b.client.do("GET", "api/v3/openOrders", "", true, &orders)

	if err != nil {
		return
	}

	return
}

// Retrieve All Open Orders for a given symbol
func (b *Binance) GetOpenOrders(query OpenOrdersQuery) (orders []OrderStatus, err error) {

	err = query.ValidateOpenOrdersQuery()
	if err != nil {
		return
	}
	reqUrl := fmt.Sprintf("api/v3/openOrders?symbol=%s&recvWindow=%d", query.Symbol, query.RecvWindow)
	_, err = b.client.do("GET", reqUrl, "", true, &orders)
	if err != nil {
		return
	}

	return
}

// Get all account orders; active, canceled, or filled.
func (b *Binance) GetAllOrders(query AllOrdersQuery) (orders []OrderStatus, err error) {
	err = query.ValidateAllOrdersQuery()
	if err != nil {
		return
	}
	reqUrl := fmt.Sprintf("api/v3/allOrders?symbol=%s&recvWindow=%d&limit=%d", query.Symbol, query.RecvWindow, query.Limit)
	if query.OrderId != 0 {
		reqUrl += fmt.Sprintf("&orderId=%d", query.OrderId)
	}
	_, err = b.client.do("GET", reqUrl, "", true, &orders)
	if err != nil {
		return
	}

	return
}

// Retrieves all trades
func (b *Binance) GetTrades(symbol string) (trades []Trade, err error) {

	reqUrl := fmt.Sprintf("api/v3/myTrades?symbol=%s", symbol)

	_, err = b.client.do("GET", reqUrl, "", true, &trades)

	if err != nil {
		return
	}
	return
}

func (b *Binance) GetTradesFromOrder(symbol string, id int64) (matchingTrades []Trade, err error) {

    reqUrl := fmt.Sprintf("api/v3/myTrades?symbol=%s", symbol)

    var trades []Trade
    _, err = b.client.do("GET", reqUrl, "", true, &trades)
    if err != nil {
        return
    }

    for _, t := range trades {
        if t.OrderId == id {
            matchingTrades = append(matchingTrades, t)
        }
    }
    return
}

//
// Retrieves all withdrawals
func (b *Binance) GetWithdrawHistory() (withdraws WithdrawList, err error) {

	reqUrl := fmt.Sprintf("wapi/v3/withdrawHistory.html")

	_, err = b.client.do("GET", reqUrl, "", true, &withdraws)
	if err != nil {
		return
	}
	return
}

//
// Retrieves all deposits
func (b *Binance) GetDepositHistory() (deposits DepositList, err error) {

	reqUrl := fmt.Sprintf("wapi/v3/depositHistory.html")

	_, err = b.client.do("GET", reqUrl, "", true, &deposits)
	if err != nil {
		return
	}
	return
}
