package profit

import "errors"

type Transaction struct {
	Amount int64
	Day    int64
}

type Profit struct {
	Buy    Transaction
	Sell   Transaction
	Profit int64
}

var PRICES_LENGTH = 0

func CalcProfit(profit *Profit, prices []int64) {

	if len(prices) == 0 {
		return
	}

	var buyPrice int64

	for day, price := range prices {
		if day == 0 {
			buyPrice = price
			continue
		}
		if (price - buyPrice) > profit.Profit {
			profit.Profit = price - buyPrice
			profit.Sell = Transaction{Amount: price, Day: int64(PRICES_LENGTH-(len(prices)-day)) + 1}
			profit.Buy = Transaction{Amount: buyPrice, Day: int64(PRICES_LENGTH-len(prices)) + 1}
		}
	}

	CalcProfit(profit, prices[1:])
}

func GetProfit(prices []int64) (*Profit, error) {

	if len(prices) < 1 || len(prices) > 100000 {
		return nil, errors.New("exchange rate list length must be  1 <= len(prices) <= 1000")
	}

	PRICES_LENGTH = len(prices)

	profit := &Profit{
		Buy:    Transaction{Amount: 0, Day: 0},
		Sell:   Transaction{Amount: 0, Day: 0},
		Profit: 0,
	}

	CalcProfit(profit, prices)

	return profit, nil

}
