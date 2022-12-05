package main

import (
	"encoding/json"
	"fmt"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

func main() {

	arr := []int64{7, 6, 5, 2, 10, 7, 9}

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
