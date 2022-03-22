package main

import (
	"fmt"
  "github.com/leekchan/accounting"
)

var(
  dollarValue= 4.91
  euroValue= 5.42
  rubleValue= 0.0094
  libraValue= 6.51
)

var (
  bigger = Green
  equal = Yellow
  lower = Red
)

func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

var (
  Red     = Color("\033[1;31m%s\033[0m")
  Green   = Color("\033[1;32m%s\033[0m")
  Yellow  = Color("\033[1;33m%s\033[0m")
)

func convertion(symbol string, value float64) string{
  currency := accounting.Accounting{Symbol: symbol, Precision: 2, Thousand: ".", Decimal: ","}

  return currency.FormatMoney(value)
}

func printData(currency string, result string) string{
  show:= fmt.Sprintf("%s: %s",currency, result)
  return show
}

func checkValue(toCheck float64, currencyValue float64, currency string, symbol string){
  result := convertion(symbol, toCheck)
  fmt.Println("")
  if toCheck > currencyValue {
    fmt.Println(bigger(printData(currency, result)))
  }
  if toCheck < currencyValue {
    fmt.Println(lower(printData(currency, result)))
  }
  if toCheck == currencyValue {
    fmt.Println(equal(printData(currency, result)))
  }
}

func calcValue(value float64, toValue float64) float64{
  return value / toValue
}

func main() {

  var value float64
  fmt.Println("Add some value, ex: 345.50")
  fmt.Scanln(&value)

  checkValue(calcValue(value, dollarValue), dollarValue, "Dollar", "$ ")
  checkValue(calcValue(value, euroValue), euroValue, "Euro", "€ ")
  checkValue(calcValue(value, rubleValue), rubleValue, "Rublo", "₽ ")
  checkValue(calcValue(value, libraValue), libraValue, "Libra", "£ ")
}
