package main

import (
  "github.com/shopspring/decimal"
  "fmt"
  "github.com/toorop/go-bittrex"
)

type CurrencyPair struct {
  Name           string
  Bases          []string
  TradingPairs   []BidAsk     //btc-alt{ask:decimal,  bid:decimal}, eth-alt{ask:decimal, bid:decimal}
}

type BidAsk struct {
  Name    string
  Bid     decimal.Decimal
  Ask     decimal.Decimal
  InvBid  decimal.Decimal
  InvAsk  decimal.Decimal
}

func (c CurrencyPair) BuildTradingPairs() CurrencyPair {
  for i:=0; i < len(c.Bases); i++ {
    tradingPair := c.Bases[i] + "-" + c.Name
    b_a := new(BidAsk)
    b_a.Name = tradingPair
    c.TradingPairs = append(c.TradingPairs, *b_a)
  }
  return c
}

func (c CurrencyPair) GetTradingPairPrices(btrxAuth *bittrex.Bittrex) (CurrencyPair, error) {
  fmt.Print(c.Name + ", ")
  for i:=0; i<len(c.TradingPairs); i++ {
      tradingPair := c.TradingPairs[i].Name

      pairInfo, err := btrxAuth.GetTicker(tradingPair)
      if err != nil {
        fmt.Println("ERROR:", err)
        return c, err
      }
      c.TradingPairs[i].Bid = pairInfo.Bid
      c.TradingPairs[i].Ask = pairInfo.Ask
      c.TradingPairs[i].InvBid, err = inverse(pairInfo.Bid)
      if err != nil {
        fmt.Println("ERROR:", err)
        return c, err
      }
      c.TradingPairs[i].InvAsk, err= inverse(pairInfo.Ask)
      if err != nil {
        fmt.Println("ERROR:", err)
        return c, err
      }
    }
    // fmt.Println(c)
  return c, nil
}

func (c CurrencyPair) FindArbitrageOpps(btrx *bittrex.Bittrex) (*TradeOrder, error) {
  UsdtExists := false
  errorTO := new(TradeOrder)
  _ = UsdtExists
  _ = errorTO

  var btc_ethBid decimal.Decimal
  var btc_ethAsk decimal.Decimal
  var btc_altBid decimal.Decimal
  var btc_altAsk decimal.Decimal
  var btc_usdBid decimal.Decimal
  var btc_usdAsk decimal.Decimal
  _, _, _, _, _, _ = btc_usdAsk, btc_usdBid, btc_altAsk, btc_altBid, btc_ethAsk, btc_ethBid

  var usd_ethAsk decimal.Decimal
  var usd_ethBid decimal.Decimal
  var usd_btcAsk decimal.Decimal
  var usd_btcBid decimal.Decimal
  var usd_altAsk decimal.Decimal
  var usd_altBid decimal.Decimal
  _, _, _, _, _, _ = usd_altBid, usd_altAsk, usd_btcBid, usd_btcAsk, usd_ethBid, usd_ethAsk

  var alt_btcAsk decimal.Decimal
  var alt_btcBid decimal.Decimal
  var alt_ethBid decimal.Decimal
  var alt_ethAsk decimal.Decimal
  var alt_usdBid decimal.Decimal
  var alt_usdAsk decimal.Decimal
  _, _, _, _, _, _ = alt_usdAsk, alt_usdBid, alt_ethAsk, alt_ethBid, alt_btcBid, alt_btcAsk

  var eth_altAsk decimal.Decimal
  var eth_altBid decimal.Decimal
  var eth_btcAsk decimal.Decimal
  var eth_btcBid decimal.Decimal
  var eth_usdAsk decimal.Decimal
  var eth_usdBid decimal.Decimal
  _, _, _, _, _, _ = eth_usdBid, eth_usdAsk, eth_btcBid, eth_btcAsk, eth_altBid, eth_altAsk

  var btc_alt_eth_btcArb decimal.Decimal
  var btc_eth_alt_btcArb decimal.Decimal
  var eth_alt_btc_ethArb decimal.Decimal
  var eth_btc_alt_ethArb decimal.Decimal
  var usd_alt_eth_usdArb decimal.Decimal
  var usd_eth_alt_usdArb decimal.Decimal
  var usd_alt_btc_usdArb decimal.Decimal
  var usd_btc_alt_usdArb decimal.Decimal
  _, _, _, _, _, _, _, _ = usd_btc_alt_usdArb, usd_alt_btc_usdArb, usd_eth_alt_usdArb, usd_alt_eth_usdArb, eth_btc_alt_ethArb, eth_alt_btc_ethArb, btc_eth_alt_btcArb, btc_alt_eth_btcArb


  for i:=0;i<len(c.Bases);i++ {
    if c.Bases[i] == "ETH" || c.Bases[i] == "BTC" {
        btc_ethTicker, err := btrx.GetTicker("BTC-ETH")
        if err != nil {
          fmt.Println("ERROR:", err)
          return errorTO, err
        }
        btc_ethBid = btc_ethTicker.Bid
        btc_ethAsk = btc_ethTicker.Ask
        eth_btcBid, err = inverse(btc_ethBid)
        if err != nil {
          fmt.Println("ERROR:", err)
          return errorTO, err
        }
        eth_btcAsk, err = inverse(btc_ethAsk)
        if err != nil {
          fmt.Println("ERROR:", err)
          return errorTO, err
        }
    } else if c.Bases[i] == "USDT" {
      UsdtExists = true
      btc_usdtTicker, err := btrx.GetTicker("USDT-BTC")
      if err != nil {
        fmt.Println("ERROR:",err)
        return errorTO, err
      }
      usd_btcAsk = btc_usdtTicker.Ask
      usd_btcBid = btc_usdtTicker.Bid
      btc_usdAsk, err = inverse(usd_btcAsk)
      if err != nil {
        fmt.Println("ERROR:", err)
        return errorTO, err
      }
      btc_usdBid, err = inverse(usd_btcBid)
      if err != nil {
        fmt.Println("ERROR:", err)
        return errorTO, err
      }

      eth_usdtTicker, err := btrx.GetTicker("USDT-ETH")
      if err != nil {
        fmt.Println("ERROR:", err)
        return errorTO, err
      }
      usd_ethAsk = eth_usdtTicker.Ask
      usd_ethBid = eth_usdtTicker.Bid
      eth_usdAsk, err = inverse(usd_ethAsk)
      if err != nil {
        fmt.Println("ERROR:", err)
        return errorTO, err
      }
      eth_usdBid,err = inverse(usd_ethBid)
      if err != nil {
        fmt.Println("ERROR:", err)
        return errorTO, err
      }
    }
  }

  for i:=0; i < len(c.TradingPairs);i++ {
    // fmt.Println(c.TradingPairs[i].Name[:3])
    if c.TradingPairs[i].Name[:3] == "BTC" {
      btc_altAsk = c.TradingPairs[i].Ask
      btc_altBid = c.TradingPairs[i].Bid
      alt_btcAsk = c.TradingPairs[i].InvAsk
      alt_btcBid = c.TradingPairs[i].InvBid
    } else if c.TradingPairs[i].Name[:3] == "ETH" {
      eth_altAsk = c.TradingPairs[i].Ask
      eth_altBid = c.TradingPairs[i].Bid
      alt_ethAsk = c.TradingPairs[i].InvAsk
      alt_ethBid = c.TradingPairs[i].InvBid
    } else if c.TradingPairs[i].Name[:3] == "USD" {
      usd_altAsk = c.TradingPairs[i].Ask
      usd_altBid = c.TradingPairs[i].Bid
      alt_usdAsk = c.TradingPairs[i].InvAsk
      alt_usdBid = c.TradingPairs[i].InvBid
    } else {
    }
  }

  per := .0025
  num := 3.0

  btc_alt_eth_btcArb = fees(alt_btcBid.Mul(eth_altAsk.Mul(btc_ethAsk)), per, num)
  BtcAltTO := new(TradeOrder)
  BtcAltTO.Base = "BTC"
  BtcAltTO.TheoreticalGain = btc_alt_eth_btcArb
  BtcAltTO.Alt = c.Name
  BtcAltTO.Trades = append(BtcAltTO.Trades, Trade{"BTC", BtcAltTO.Alt, "Bid", "BTC-" + BtcAltTO.Alt})
  BtcAltTO.Trades = append(BtcAltTO.Trades, Trade{"ETH", BtcAltTO.Alt, "Ask", "ETH-" + BtcAltTO.Alt})
  BtcAltTO.Trades = append(BtcAltTO.Trades, Trade{"BTC", "ETH", "Ask", "BTC-ETH"})

  btc_eth_alt_btcArb = fees(eth_btcBid.Mul(alt_ethBid.Mul(btc_altAsk)), per, num)
  BtcEthTO := new(TradeOrder)
  BtcEthTO.Base = "BTC"
  BtcEthTO.TheoreticalGain = btc_eth_alt_btcArb
  BtcEthTO.Alt = c.Name
  BtcEthTO.Trades = append(BtcEthTO.Trades, Trade{"BTC", "ETH", "Bid", "BTC-ETH"})
  BtcEthTO.Trades = append(BtcEthTO.Trades, Trade{"ETH", BtcAltTO.Alt, "Bid", "ETH-" + BtcEthTO.Alt})
  BtcEthTO.Trades = append(BtcEthTO.Trades, Trade{"BTC", BtcAltTO.Alt, "Ask", "BTC-" + BtcEthTO.Alt})

  if BtcEthTO.TheoreticalGain.GreaterThanOrEqual(BtcAltTO.TheoreticalGain) {
    return BtcEthTO, nil
  }
  return BtcAltTO, nil

  // eth_alt_btc_ethArb = alt_ethBid.Mul(btc_altAsk.Mul(eth_btcAsk))
  // eth_btc_alt_ethArb = btc_ethAsk.Mul(alt_btcBid.Mul(eth_altAsk))

  // if UsdtExists {
  //   usd_alt_eth_usdArb = alt_usdBid.Mul(eth_altAsk.Mul(usd_ethAsk))
  //   usd_eth_alt_usdArb = eth_usdAsk.Mul(alt_ethBid.Mul(usd_altAsk))
  //
  //   usd_alt_btc_usdArb = alt_usdBid.Mul(btc_altAsk.Mul(usd_btcAsk))
  //   usd_btc_alt_usdArb = btc_usdAsk.Mul(alt_btcBid.Mul(usd_altAsk))
  // }

  // baeb := fees(btc_alt_eth_btcArb, per, num))
  // beab := fees(btc_eth_alt_btcArb, per, num)
  // eabe := fees(eth_alt_btc_ethArb, per, num)
  // ebae := fees(eth_btc_alt_ethArb, per, num)
  // fmt.Println(usd_alt_eth_usdArb)
  // fmt.Println(usd_eth_alt_usdArb)
  // fmt.Println(usd_alt_btc_usdArb)
  // fmt.Println(usd_btc_alt_usdArb)
  // usdt_btcTicker, err := btrx.GetTicker("USDT-BTC")
  // uTicker, err := btrx.GetTicker("USDT-ETH")
}

func (c CurrencyPair) AddBase(base string) CurrencyPair {
  // fmt.Println(c.Name)
  baseExists := false
  for i := 0; i < len(c.Bases); i++ {
    if base == c.Bases[i] {
      baseExists = true
      break
    } else {
      baseExists = false
    }
  }
  if !baseExists {
    c.Bases = append(c.Bases, base)
    // fmt.Println(c.Bases)
  }
  return c
}
