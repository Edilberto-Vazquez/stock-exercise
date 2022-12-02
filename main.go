package main

import (
	"encoding/json"
	"fmt"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

func main() {
	actions := [][]int64{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{7, 10001, 1, 3, 1},
		{},
		make([]int64, 100001),
	}

	for _, action := range actions {
		fmt.Println("**********")
		if len(action) > 100000 {
			fmt.Println("Input: [0,0,0,...0]")
		} else {
			fmt.Printf("Input: %v\n", action)
		}

		result, err := profit.CalcProfit(action)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			res, _ := json.Marshal(result)
			fmt.Printf("%s\n", string(res))
		}

		fmt.Println("**********")
	}

}
