package main

import (
	"encoding/json"
	"fmt"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

func main() {

	arr := []int64{7, 4, 2, 10000, 6, 3, 8, 5}

	fmt.Println("**********")

	fmt.Println(arr)

	profit, err := profit.GetProfit(arr)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		res, _ := json.Marshal(*profit)
		fmt.Printf("%s\n", string(res))
	}

	fmt.Println("**********")

}
