package main

import (
  "fmt"
  "github.com/toorop/go-bittrex"
  "bufio"
  "os"
  "github.com/shopspring/decimal"
  "errors"
  "time"
)

func authentication() (string, string) {
  fmt.Println("Authenticating...")
  inFile, _ := os.Open("./config")
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
  scanner.Scan()
  scanner.Scan()
  scanner.Scan()
  API_KEY := scanner.Text()                                                     //Line 3
  scanner.Scan()
  scanner.Scan()
  API_SECRET := scanner.Text()                                                  //Line 5
  inFile.Close()
  return API_KEY, API_SECRET
}

func inverse(num decimal.Decimal) (decimal.Decimal, error) {
  // fmt.Println(num)
  if num == decimal.NewFromFloat(0) {
    return decimal.NewFromFloat(0), errors.New("Divide by 0")
  }
  if num != decimal.NewFromFloat(0) {
    return decimal.NewFromFloat(1.00).DivRound(num, 10), nil
  }
  _ = time.Second
  return decimal.NewFromFloat(0.00), errors.New("Divide by 0")
}

func fees(num decimal.Decimal, percent float64, numberOfTrades float64) decimal.Decimal {
  fee := decimal.NewFromFloat(1.0).Sub(decimal.NewFromFloat(percent*numberOfTrades))
  num = num.Mul(fee)
  // fmt.Println(fee)
  return num
}

func performArbitrage(tradeInfo *TradeOrder, btrx *bittrex.Bittrex, fakeInfo FalseTrade) (History, error, FalseTrade) {
  var price decimal.Decimal
  var amountToTrade decimal.Decimal
  var hist History
  _ = amountToTrade

  if !RealTrading {
    hist.StartAmount = fakeInfo.CurrentAmount
    hist.TheoreticalPer = tradeInfo.TheoreticalGain.Sub(decimal.NewFromFloat(1))
  }

  for i:=0;i<len(tradeInfo.Trades);i++ {
    side := tradeInfo.Trades[i].BidAsk

    pairInfo, err := btrx.GetTicker(tradeInfo.Trades[i].Pair)
    if err != nil {
      fmt.Println("ERROR performAribtrage: ", err)
      return hist, errors.New("Could not find ticker"), fakeInfo
    }
    if !RealTrading {         //Write fake trading junk here

      if side == "Bid" {
        price = pairInfo.Bid
        amountToTrade, fakeInfo, err = getAmountToTrade(tradeInfo.Trades[i], price, fakeInfo)
        if err != nil {
          return hist, err, fakeInfo
        }
      } else {
        price = pairInfo.Ask
        amountToTrade, fakeInfo, err = getAmountToTrade(tradeInfo.Trades[i], price, fakeInfo)
        if err != nil {
          return hist, err, fakeInfo
        }
      }

      time.Sleep(5 * time.Second)
    } else if RealTrading {   //Write real trading junk here
      if side == "Bid" {
        price = pairInfo.Bid
        amountToTrade, fakeInfo, err = getAmountToTrade(tradeInfo.Trades[i], price, fakeInfo)
        // uuid, err := btrx.BuyLimit(tradeInfo.Trades[i].Pair, amountToTrade, price)
      } else {
        price = pairInfo.Ask
        amountToTrade, fakeInfo, err = getAmountToTrade(tradeInfo.Trades[i], price, fakeInfo)
        // uuid, err := bittrex.SellLimit(tradeInfo[i].Pair, amountToTrade, price)
      }
    }
  }

  if !RealTrading {
    hist.EndAmount = fakeInfo.CurrentAmount
    hist.PercentGain = hist.EndAmount.Sub(hist.StartAmount).Div(hist.StartAmount)
    hist.PairTraded = tradeInfo.Alt
    hist.Time = time.Now()
  }

  return hist, nil, fakeInfo
}

func getAmountToTrade(trade Trade, price decimal.Decimal, fakeInfo FalseTrade) (decimal.Decimal, FalseTrade, error) {
  if !RealTrading {
    per := .0025
    num := 1.0
    // fmt.Println("ALT:",trade.Alt)
    // fmt.Println("BASE:",trade.Base)
    if price == decimal.NewFromFloat(0.0) || fakeInfo.CurrentAmount == decimal.NewFromFloat(0.0) {
      return decimal.NewFromFloat(0.0), fakeInfo, errors.New("Price Cant Be Zero")
    }
    // fakeInfo.StartAmount = fakeInfo.CurrentAmount
    // fmt.Println("fake",fakeInfo)
    if trade.BidAsk == "Bid" && trade.Base == fakeInfo.CurrentHolding {
      amount := fakeInfo.CurrentAmount.DivRound(price, 10)
      fakeInfo.CurrentHolding = trade.Alt
      fakeInfo.CurrentAmount = fees(amount, per, num)
      return amount, fakeInfo, nil
    } else if trade.BidAsk == "Ask" && trade.Alt == fakeInfo.CurrentHolding {
      amount := fakeInfo.CurrentAmount.Mul(price)
      fakeInfo.CurrentHolding = trade.Base
      fakeInfo.CurrentAmount = fees(amount, per, num)
      return amount, fakeInfo, nil
    }
  }

  return decimal.NewFromFloat(0.0), fakeInfo, nil
}
