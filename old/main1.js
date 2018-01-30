const key          = '' // API Key
const secret       = '' // API Private Key
const KrakenClient = require('kraken-api')
var sleep = require('sleep')
const kraken       = new KrakenClient(key, secret)

// Get Ticker Info
ethbtc = kraken.api('Ticker', {"pair": 'XETHXXBT'}, async function(error, data) {
    if(error) {
        console.log(error)
    }
    else {
        // console.log(data.result)
        return await data.result['XETHXXBT']['a'][0]
    }

})

etceth = kraken.api('Ticker', {"pair": 'XETCXETH'}, async function(error, data) {
    if(error) {
        console.log(error)
    }
    else {
        // console.log(data.result)
        return await data.result['XETCXETH']['a'][0]
    }
    // console.log(etceth)
})

etxbtc = kraken.api('Ticker', {"pair": 'XETCXXBT'}, async function(error, data) {
    if(error) {
        console.log(error)
    }
    else {
        // console.log(data.result)
        return await data.result['XETCXXBT']['a'][0]
    }
    // console.log(etcbtc)
})

var etc_btc = etceth*ethbtc
console.log(etc_btc)

// kraken.api('AddOrder', {"pair": 'XXBTZUSD', "type": 'sell', "ordertype": 'market', "volume": .01}, function(error, data) {
//     if(error) {
//         console.log(error);
//     }
//     else {
//         console.log(data.result);
//     }
// })

// Display user's balance
// kraken.api('Balance', null, function(error, data) {
//     if(error) {
//         console.log(error);
//     }
//     else {
//       // console.log('T')
//       console.log(data.result['ZUSD']);
//     }
// })
