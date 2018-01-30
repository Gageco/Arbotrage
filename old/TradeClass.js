class Trade {
  constructor(to, from, coin, amount, percent) {
    this.To = to                  //What exchange the coin is to go to
    this.From = from              //What exchange the coin is coming from
    this.Coin = coin              //What coin is being transfered between exchanges
    this.Amount = amount          //The amount of coin that is going to be transfered
    this.Percent = percent        //The percentage that was calculated
    //I think i can do something with this.Percent where it checks the percent and checks liquidity or something to see if i need to move money or not

    // this.Action = action


    /*

    yo future gage figure out how to do default constructors so then you wont have to have undefined stuff

    */
  }
}

module.exports = Trade
