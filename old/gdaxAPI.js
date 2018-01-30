const Gdax = require('gdax')
const pubClient = new Gdax.PublicClient();

class GdaxAP {

  async getPrice(coin) {

    const myCallback = async (err, response, data) => {
      if (err == null) {
        // console.log(await data['price'])
        let x = await data['price']
        console.log(await x)
        return x
      } else {
        console.log("Error: " + err)
      }
    }

    var coinTrade = coin + '-usd'
    pubClient.getProductTicker(coinTrade, myCallback)
      // .then(data => {
      //   // console.log(data['price'])
      //   return data['price']
      // })

  }
}


module.exports = GdaxAP
