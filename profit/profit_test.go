package profit_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Edilberto-Vazquez/stock-exercise/profit"
)

type Test struct {
	Message string
	Input   []int64
	Result  *profit.Profit
	Err     error
}

func TestCalcProfit(t *testing.T) {
	t.Log("Given an array of prices for day")
	t.Log("When it is required to maximize profit by choosing a single day to buy and a different day in the future to sell")
	t.Logf("Note that you cannot buy on day 2 and sell on day 1. You have to buy before you sell.\n\n")

	testsInputsOk := []Test{
		{
			Message: "Then returns the maximum profit for the transaction and a nil error",
			Input:   []int64{7, 1, 5, 3, 6, 4},
			Result: &profit.Profit{
				Buy:    profit.Transaction{Amount: 1, Day: 2},
				Sell:   profit.Transaction{Amount: 6, Day: 5},
				Profit: 5,
			},
		},
		{
			Message: "Then if you cant get any profit return 0 and a nil error",
			Input:   []int64{7, 6, 4, 3, 1},
			Result: &profit.Profit{
				Buy:    profit.Transaction{Amount: 1, Day: 5},
				Sell:   profit.Transaction{Amount: 0, Day: 0},
				Profit: 0,
			},
		},
		{
			Message: "Then if some price is less than or equal 10000 returns the maximum and a nil error",
			Input:   []int64{7, 4, 2, 10000, 6, 3, 8, 5},
			Result: &profit.Profit{
				Buy:    profit.Transaction{Amount: 2, Day: 3},
				Sell:   profit.Transaction{Amount: 10000, Day: 4},
				Profit: 9998,
			},
		},
		{
			Message: "Then if some price is greater than 10000 skip value returns the maximum and a nil error",
			Input:   []int64{7, 4, 2, 10001, 6, 3, 8, 5},
			Result: &profit.Profit{
				Buy:    profit.Transaction{Amount: 2, Day: 3},
				Sell:   profit.Transaction{Amount: 8, Day: 7},
				Profit: 6,
			},
		},
	}

	testsInputsWrong := []Test{
		{
			Message: "Then if the array is empty return a nil pointer and an error",
			Input:   []int64{},
			Result:  nil,
		},
		{
			Message: "Then if the array is greater than 100000 return a nil pointer and an error",
			Input:   make([]int64, 100001),
			Result:  nil,
		},
	}

	for _, test := range testsInputsOk {
		t.Log(test.Message)
		result, _ := profit.CalcProfit(test.Input)
		testRes, _ := json.Marshal(test.Result)
		resultRes, _ := json.Marshal(result)
		if reflect.DeepEqual(test.Result, result) {
			t.Logf("CalcProfit(%v) PASSED.", test.Input)
			t.Logf("Expected %s", testRes)
			t.Logf("Got %s\n\n", resultRes)
		} else {
			t.Errorf("CalcProfit(%v) FAILED.", test.Input)
			t.Errorf("Expected %s", testRes)
			t.Errorf("Got %s\n\n", resultRes)
		}
	}

	for _, test := range testsInputsWrong {
		t.Log(test.Message)
		result, err := profit.CalcProfit(test.Input)
		if err != nil && result == nil {
			if len(test.Input) > 100000 {
				t.Log("CalcProfit([0,0,0,...0]) PASSED.")
				t.Logf("Expected %v", test.Result)
				t.Logf("Got %v\n\n", result)
			} else {
				t.Logf("CalcProfit(%v) PASSED.", test.Input)
				t.Logf("Expected %v", test.Result)
				t.Logf("Got %v\n\n", result)
			}
		} else {
			if len(test.Input) > 100000 {
				t.Log("CalcProfit([0,0,0,...0]) FAILED.")
				t.Logf("Expected %v", test.Result)
				t.Logf("Got %v\n\n", result)
			} else {
				t.Logf("CalcProfit(%v) FAILED.", test.Input)
				t.Logf("Expected %v", test.Result)
				t.Logf("Got %v\n\n", result)
			}
		}
	}
}
