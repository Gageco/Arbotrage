package main
// Get Ticker (BTC-VTC)

// ticker, err := bittrex.GetTicker("BTC-LTC")//BTC/LTC
// fmt.Println(err, ticker)
//
// ticker1, err := bittrex.GetTicker("ETH-LTC")//ETH/LTC
// fmt.Println(err, ticker1)
//
// ticker2, err := bittrex.GetTicker("BTC-ETH")//BTC/ETH
// fmt.Println(err, ticker2)
//
// ltc_btc := decimal.NewFromFloat(1.000).DivRound(ticker.Bid, 10)
// eth_ltc := ticker1.Bid
// btc_eth := ticker2.Bid
//
// eth_btc := decimal.NewFromFloat(1.000).DivRound(btc_eth, 10)
// ltc_eth := decimal.NewFromFloat(1.000).DivRound(eth_ltc, 10)
// btc_ltc := decimal.NewFromFloat(1.000).DivRound(ltc_btc, 10)
// // fmt.Println(ticker.Ask)
// x := btc_eth.Mul(ltc_btc.Mul(eth_ltc))
// y := eth_btc.Mul(btc_ltc.Mul(ltc_eth))
// fmt.Println(x)
// fmt.Println(y.Sub(y.Mul(decimal.NewFromFloat(.0025*3))))

// // Get Distribution (JBS)
// 	distribution, err := bittrex.GetDistribution("JBS")
// 	for _, balance := range distribution.Distribution {
// 		fmt.Println(balance.BalanceD)
// 	}
//
// // Get market summaries
// 	marketSummaries, err := bittrex.GetMarketSummaries()
// 	fmt.Println(err, marketSummaries)
//
// // Get market summary
// 	marketSummary, err := bittrex.GetMarketSummary("BTC-ETH")
// 	fmt.Println(err, marketSummary)
//
// // Get orders book
//
// 	orderBook, err := bittrex.GetOrderBook("BTC-DRK", "both")
// 	fmt.Println(err, orderBook)
//
//
// // Get order book buy or sell side only
//
// 	orderb, err := bittrex.GetOrderBookBuySell("BTC-JBS", "buy")
// 	fmt.Println(err, orderb)
//
//
// // Market history
//
// 	marketHistory, err := bittrex.GetMarketHistory("BTC-DRK")
// 	for _, trade := range marketHistory {
// 		fmt.Println(err, trade.Timestamp.String(), trade.Quantity, trade.Price)
// 	}
//
//
// // Market
//
// // BuyLimit
//
// 	uuid, err := bittrex.BuyLimit("BTC-DOGE", 1000, 0.00000102)
// 	fmt.Println(err, uuid)
//
//
// // Sell limit
//
// 	uuid, err := bittrex.SellLimit("BTC-DOGE", 1000, 0.00000115)
// 	fmt.Println(err, uuid)
//
//
// // Cancel Order
//
// 	err := bittrex.CancelOrder("e3b4b704-2aca-4b8c-8272-50fada7de474")
// 	fmt.Println(err)
//
//
// // Get open orders
//
// 	orders, err := bittrex.GetOpenOrders("BTC-DOGE")
// 	fmt.Println(err, orders)
//
//
// // Account
// // Get balances
//
// 	balances, err := bittrex.GetBalances()
// 	fmt.Println(err, balances)
//
//
// // Get balance
//
// 	balance, err := bittrex.GetBalance("DOGE")
// 	fmt.Println(err, balance)
//
//
// // Get address
//
// 	address, err := bittrex.GetDepositAddress("QBC")
// 	fmt.Println(err, address)
//
//
// // WithDraw
//
// 	whitdrawUuid, err := bittrex.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
// 	fmt.Println(err, whitdrawUuid)
//
//
// // Get order history
//
// 	orderHistory, err := bittrex.GetOrderHistory("BTC-DOGE")
// 	fmt.Println(err, orderHistory)
//
//
// // Get getwithdrawal history
//
// 	withdrawalHistory, err := bittrex.GetWithdrawalHistory("all")
// 	fmt.Println(err, withdrawalHistory)
//
//
// // Get deposit history
//
// 	deposits, err := bittrex.GetDepositHistory("all")
// 	fmt.Println(err, deposits)
