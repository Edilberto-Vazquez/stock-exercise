package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	Amount int64
	Day    int64
}

type Profit struct {
	Buy    Transaction
	Sell   Transaction
	Profit int64
}

func CalcProfit(exchangeRates []int64) (*Profit, error) {

	if len(exchangeRates) == 0 {
		return nil, errors.New("exchange rate list length is too short must be equal or greater than 1")
	}

	if len(exchangeRates) > 100000 {
		return nil, errors.New("the length of the exchange rate list is too long")
	}

	profit := &Profit{
		Buy:    Transaction{Amount: 0, Day: 0},
		Sell:   Transaction{Amount: 0, Day: 0},
		Profit: 0,
	}

	for day, amount := range exchangeRates {
		if amount <= 0 || amount > 10000 {
			continue
		}

		if amount <= profit.Buy.Amount || profit.Buy.Amount == 0 {
			profit.Buy.Amount = amount
			profit.Buy.Day = int64(day + 1)
		}

		if day > 0 && day > int(profit.Buy.Day-1) && amount > profit.Buy.Amount && amount > profit.Sell.Amount {
			profit.Sell.Amount = amount
			profit.Sell.Day = int64(day + 1)
		}
	}

	if profit.Sell.Amount-profit.Buy.Amount > 0 {
		profit.Profit = profit.Sell.Amount - profit.Buy.Amount
		return profit, nil
	}

	return profit, nil

}

func main() {

	actions1 := []int64{7, 1, 5, 3, 6, 4}
	actions2 := []int64{7, 6, 4, 3, 1}
	actions3 := []int64{7, 10001, 1, 3, 1}
	actions4 := []int64{}
	actions5 := make([]int64, 100001)

	actions := [][]int64{actions1, actions2, actions3, actions4, actions5}

	for _, action := range actions {
		fmt.Println("**********")
		fmt.Printf("Input: %v\n", action)

		profit, err := CalcProfit(action)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Buy day: %d, Amount: %d;\n", profit.Buy.Day, profit.Buy.Amount)
			fmt.Printf("Sell day: %d, Amount: %d;\n", profit.Sell.Day, profit.Sell.Amount)
			fmt.Printf("Profit: %d\n", profit.Profit)
		}

		fmt.Println("**********")
	}

}
