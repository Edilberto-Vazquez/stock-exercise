package profit

import (
	"errors"
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

func GetProfit(prices []int64) (*Profit, error) {

	if len(prices) < 1 || len(prices) > 100000 {
		return nil, errors.New("exchange rate list length must be  1 <= len(prices) <= 1000")
	}

	profit := Profit{
		Buy:    Transaction{Amount: 0, Day: 0},
		Sell:   Transaction{Amount: 0, Day: 0},
		Profit: 0,
	}

	for day, price := range prices {
		if day == 0 {
			continue
		}
		if price > profit.Sell.Amount {
			profit.Sell = Transaction{Amount: price, Day: int64(day) + 1}
		}
	}

	for day, price := range prices {
		if day+1 == int(profit.Sell.Day) {
			continue
		}
		if day+1 < int(profit.Sell.Day) && price < profit.Sell.Amount {
			if profit.Sell.Amount-price > profit.Profit {
				profit.Profit = profit.Sell.Amount - price
				profit.Buy = Transaction{Amount: price, Day: int64(day) + 1}
			}
		}
	}

	return &profit, nil

}
