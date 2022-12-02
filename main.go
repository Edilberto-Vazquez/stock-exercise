package main

import (
	"fmt"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

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

		result, err := profit.CalcProfit(action)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Buy day: %d, Amount: %d;\n", result.Buy.Day, result.Buy.Amount)
			fmt.Printf("Sell day: %d, Amount: %d;\n", result.Sell.Day, result.Sell.Amount)
			fmt.Printf("Profit: %d\n", result.Profit)
		}

		fmt.Println("**********")
	}

}
