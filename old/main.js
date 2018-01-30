const Exchange = require('./ExchangeClass.js')
const Trade = require('./TradeClass.js')

var kraken = new Exchange('kraken')
var gdax = new Exchange('gdax')
var bittrex = new Exchange('bittrex')


var coinList = ['btc', 'eth'] //supported currencies
var exchanges = {'kraken': kraken, 'gdax': gdax, 'bittrex': bittrex}  //supported exchanges

initilization()
updateExchInfo()
// console.log("1")


// exchanges['kraken'].Amount['btc'] = 3
//
// exchanges['kraken'].Price['btc'] = 13000
// exchanges['gdax'].Price['btc'] = 15000
// exchanges['bittrex'].Price['btc'] = 14000
//
// calculateExchDiff()
// // console.log(exchanges['gdax'])
//
// trades = determineTrades()
// performTrades()

function performTrades() {
  for(t in trades) {
    fromExch = exchanges[trades[t].From]    //Get the exchange from the trade
    toExch = exchanges[trades[t].To]         //Get the exchange to transfer to from the trade
    coin = trades[t].Coin
    // console.log(fromExch)
    if (fromExch != undefined) {
      if (fromExch.Amount[coin] > trades[t].Amount) {
        address = toExch.getAddress(coin)
        fromExch.transferTo(address, coin)
        // console.log(trades[t])
        // updateExchInfo(exchanges)
        console.log("Arbitrage: " + coin +" From: " + fromExch.Name + ": " + " " + fromExch.Price[coin] + ", To: " + toExch.Name + " " + toExch.Price[coin] + " for $" + (toExch.Price[coin] - fromExch.Price[coin]) + ' profit')

      } else {
        console.log("Not Enough in " + fromExch.Name + " to send " + trades[t].Amount + " " + coin + " to " + toExch.Name)
      }
    }
  }
}

function determineTrades() {
  var trades = [0, 0, 0, 0]

  for (exch in exchanges) {
    for (exch0 in exchanges) {
      if (exch != exch0) {
        for (coin in coinList) {
          var sameExchange = false                                                //Checking if it is the same exchange
          var percChange = exchanges[exch].ExchangeDiff[exch0][coinList[coin]]

          if (percChange >= 0.05) {                              //If there is a 5% or greater difference between exchanges
            // console.log("New High Trade: " + percChange)
            var newTrade = new Trade(exch, exch0, coinList[coin], 1, percChange) //NOTE CHANGE THIS VALUE
            for (i in trades) {
              if (trades[i] == 0) {
                trades[i] = newTrade
                break
              } else if (i == trades.length) {
                // console.log("No more room for trades: " + percChange)
              }
            }
          } else if (percChange <= 0.01 && percChange > 0.00) {                       //If there is a 1% or less difference between exchanges
            // console.log("New Low Trade" + percChange)
            var newTrade = new Trade(exch, exch0, coinList[coin], 1, percChange)
            for (i in trades) {
              if (trades[i] == 0) {
                trades[i] = newTrade
              } else if (i == trades.length) {
                console.log("No more room for trades")
              }
            }
          }
        }
      }
    }
  }
  // console.log(trades)
  return trades

}

async function updateExchInfo() {
  for (exch in exchanges) {
    for (coin in coinList) {
      await exchanges[exch].updatePrices(coinList[coin])
      await exchanges[exch].updateAmount(coinList[coin])
    }
  }
  console.log(await "UEF")
  console.log(await exchanges['gdax'].Price)
}

function calculateExchDiff() {
  for (i = 0; i < Object.keys(exchanges).length; i++) {
    exch = Object.keys(exchanges)[i]
    for (j = 0; j < Object.keys(exchanges).length; j++) { //Object.keys(exchanges).length
      exch0 = Object.keys(exchanges)[j]
      if (exch != exch0) {
        for (k = 0; k < coinList.length; k++) {
          exchangePrice00 = exchanges[exch].Price[coinList[k]]
          exchangePrice01 = exchanges[exch0].Price[coinList[k]]

          var percChange = (exchangePrice00 - exchangePrice01)/exchangePrice01
          // // console.log("CEF: "+ exch + "/" + exch0 + " " + percChange)
          // // console.log("BEFORE")
          // console.log(exchanges[exch].ExchangeDiff)

          exchanges[exch].ExchangeDiff[exch0][coinList[k]] = percChange
          // console.log("AFTER")
          // console.log(exchanges[exch].ExchangeDiff)//[exch0][coinList[coin]])
            // console.log(exchanges[exch])
        }
      }
    }
  }
}


function initilization() {        //learn how to pass by reference and move to outside file, pass exchangePrices and exchangePriceDifference
  for (exch in exchanges) {

    for (coin in coinList) {
      exchanges[exch].Amount[coinList[coin]] = 0.000001
      exchanges[exch].Price[coinList[coin]] = 0.000001
      for (exch0 in exchanges) {
        if (exch != exch0) {
          // console.log("Init")
          exchanges[exch].ExchangeDiff[exch0] = {}
          exchanges[exch].ExchangeDiff[exch0][coinList[coin]] = 0.000001
          // console.log(exchanges[exch].ExchangeDiff[exch0])
        }
      }
    }
  }
}
