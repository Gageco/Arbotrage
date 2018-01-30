const GdaxAPI = require('./gdaxAPI.js')

class Exchange {
  constructor(name) {
    this.Name = name
    this.Amount = {}
    this.ExchangeDiff = {}
    this.Price = {}
    // this.fees <-- NOTE dont know what this will look like but to take fees into calculations
  }

  transferTo(address, coin) {        //for each exchange transfer to a given address
    switch (this.Name) {
      case "kraken":
        //NOTE Code for kraken transfer here
        break
      case "gdax":
        //NOTE Code for gdax transfer here
        break
      default:
        break
    }
  }

  authentication() {
    switch (this.Name) {
      case "kraken":
        //NOTE Code for kraken authentication goes here
        break
      case "gdax":
        //NOTE Code for gdax authentication goes here
        break
      default:
        break
    }
  }

  async updatePrices(coin) {
    switch (this.Name) {
      case "kraken":
        //NOTE: This is where i would update the prices for each exchannge
        break
      case "gdax":
        // console.log(await GdaxAPI.getPrice(coin))
        this.Price[coin] = await GdaxAPI.getPrice(coin)
        break
      default:
        break
    }
  }

  updateAmount() {
    switch (this.Name) {
      case "kraken":
        //NOTE Code for kraken get update on liquidity here, accept coin as variable
        break
      case "gdax":
        //NOTE Code for gdeax get update on liquidity here, accept coin as variable
        break
      default:
        break
    }
  }

  getAddress() {
    switch (this.Name) {
      case "kraken":
        //NOTE Code for kraken get an address here, accept coin as variable to get specific address
        break
      case "gdax":
        //NOTE Code for gdax get an address here, accept coin as variable to get specific address
        break
      default:
        break
    }
  }

}

module.exports = Exchange
