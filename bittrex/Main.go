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
	var fakeInfo FalseTrade
	// usdt_btcTicker, err := btrx.GetTicker("USDT-BTC")
	API_KEY, API_SECRET := authentication()
	bittrex := bittrex.New(API_KEY, API_SECRET)

	markets, err := bittrex.GetMarkets()
  if err != nil {
    fmt.Println(err)
  }

  var baseCurrency []string
  var arbitrageCurrency []string
  var tokenMap = make(map[string]CurrencyPair)

	if !RealTrading {
		fakeInfo = buildFakeTrading()
	}

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

for g:=0; g<10; g++ {
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
		if tradeInfo.TheoreticalGain.GreaterThanOrEqual(decimal.NewFromFloat(1.05)) {
			// fmt.Println(tradeInfo.TheoreticalGain)
			info, err, fakeInfo := performArbitrage(tradeInfo, bittrex, fakeInfo)
			_, _ = info, err
			fmt.Println("Traded:",tradeInfo.Alt)
			fmt.Println(fakeInfo)
		}
  }
}

  // fmt.Println(tokenMap["BNT"])
}
