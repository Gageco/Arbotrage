package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
  // "./CurrencyPair.go"
  "github.com/shopspring/decimal"
)

const (
	RealTrading bool = false
)

func main() {
	API_KEY, API_SECRET := authentication()
	bittrex := bittrex.New(API_KEY, API_SECRET)
	fmt.Println("Getting Bittrex Markets...")
	markets, err := bittrex.GetMarkets()
  if err != nil {
    fmt.Println(err)
  }

	var history             []History
	var fakeInfo            FalseTrade
  var baseCurrency        []string
  var arbitrageCurrency   []string
  var tokenMap = make(map[string]CurrencyPair)

	fakeInfo = FalseTrade{"BTC", decimal.NewFromFloat(.025),"BTC",decimal.NewFromFloat(.025)}

	fmt.Println("Sorting Through Markets...")
  for i:=0; i < len(markets); i++ {
    foundBaseCurrency := false
    foundArbitrageCurrency := false
    if i == 0 {
      baseCurrency = append(baseCurrency, markets[i].BaseCurrency)
      arbitrageCurrency = append(arbitrageCurrency, markets[i].MarketCurrency)
    }

    for j:=0; j < len(arbitrageCurrency); j++ {
      if arbitrageCurrency[j] == markets[i].MarketCurrency {
        foundArbitrageCurrency = true
        break
      } else {
        foundArbitrageCurrency = false
      }
    }
    if !foundArbitrageCurrency {
      arbitrageCurrency = append(arbitrageCurrency, markets[i].MarketCurrency)
    }

    for j:=0; j < len(baseCurrency); j++ {
      if baseCurrency[j] == markets[i].BaseCurrency {
        foundBaseCurrency = true
        break
      } else {
        foundBaseCurrency = false
      }
    }
    if !foundBaseCurrency {
      baseCurrency = append(baseCurrency, markets[i].BaseCurrency)
    }
  }

  // fmt.Println(baseCurrency)
  // fmt.Println(arbitrageCurrency)

  for i:=0; i < len(arbitrageCurrency); i++ {
    pair := CurrencyPair{}
    pair.Name = arbitrageCurrency[i]
    tokenMap[arbitrageCurrency[i]] = pair
  }

  for i:=0; i < len(baseCurrency); i++ {
    for j:=0; j < len(arbitrageCurrency); j++ {
      testMarket := baseCurrency[i] + "-" + arbitrageCurrency[j]
      for k:=0; k < len(markets); k++ {
        if testMarket == markets[k].MarketName {
          tokenMap[arbitrageCurrency[j]] = tokenMap[arbitrageCurrency[j]].AddBase(baseCurrency[i])
        }
      }
    }
  }

  for key, pair := range tokenMap {
    if len(pair.Bases) == 1 {
      delete(tokenMap, key)
    } else {
      tokenMap[key] = pair.BuildTradingPairs()
    }
  }
	fmt.Println("Determing Arbitrage Opportunities...")
	for g:=0; g<5; g++ {
		fmt.Print("Checking: ")
	  for key, pair := range tokenMap {
	    tokenMap[key], err = pair.GetTradingPairPrices(bittrex)
			if err != nil {
				fmt.Println("Err:",err)
				break
			}
			// fmt.Println(tokenMap[key])
			tradeInfo, err := pair.FindArbitrageOpps(bittrex)
			if err != nil {
				fmt.Println("Err:",err)
				break
			}
			if tradeInfo.TheoreticalGain.GreaterThanOrEqual(decimal.NewFromFloat(1.02)) {
				fmt.Println("\nOpportunity Found", tradeInfo.Alt)
				// fmt.Println(tradeInfo.TheoreticalGain)
				hist, err, FI := performArbitrage(tradeInfo, bittrex, fakeInfo)
				_ = err
				fakeInfo = FI
				history = append(history, hist)
				// fmt.Println(history)
				fmt.Println(fakeInfo)
				fmt.Print("\nChecking: ")
			}
	  }
	}
}
