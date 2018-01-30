package main

import (
  "github.com/shopspring/decimal"
  "time"
)

type TradeOrder struct {
  Base              string
  Alt               string
  Trades            []Trade
  TheoreticalGain   decimal.Decimal
}

type Trade struct {
  Base       string
  Alt        string
  BidAsk     string
  Pair       string
}

type History struct {
  Time           time.Time //Some sort of time thing
  StartAmount    decimal.Decimal
  EndAmount      decimal.Decimal
  PercentGain    decimal.Decimal
  TheoreticalPer decimal.Decimal
  PairTraded     string
}

type FalseTrade struct {
  StartHolding        string
  StartAmount         decimal.Decimal
  CurrentHolding      string
  CurrentAmount       decimal.Decimal
}
