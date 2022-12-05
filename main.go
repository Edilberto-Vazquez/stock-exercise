package main

import (
	"encoding/json"
	"fmt"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

func main() {
	actions := [][]int64{
		{7, 4, 10, 9, 3, 6, 3, 10},
		{7, 6, 5, 2, 10, 7, 9},
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{7, 4, 2, 10000, 6, 3, 8, 5},
		{7, 4, 8, 3, 2, 4, 3, 5},
		// {},
		// make([]int64, 100001),
	}

	for _, action := range actions {
		fmt.Println("**********")
		if len(action) > 100000 {
			fmt.Println("Input: [0,0,0,...0]")
		} else {
			fmt.Printf("Input: %v\n", action)
		}

		profit, err := profit.GetProfit(action)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			res, _ := json.Marshal(*profit)
			fmt.Printf("%s\n", string(res))
		}

		fmt.Println("**********")
	}

}
