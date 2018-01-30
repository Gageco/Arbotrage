package main

import (
  "github.com/shopspring/decimal"
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
  Time           string //Some sort of time thing
  StartAmount    string
  EndAmount      string
  PercentGain    decimal.Decimal
  PairTraded     string
}

type FalseTrade struct {
  StartHolding        string
  StartAmount         decimal.Decimal
  CurrentHolding      string
  CurrentAmount       decimal.Decimal
}
