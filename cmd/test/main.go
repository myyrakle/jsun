package main

import (
	"fmt"

	json "github.com/myyrakle/jsun"
)

type TradeDetail struct {
	ItemID string `json:"item_id"`
}

type TradeHistory struct {
	TradeID string        `json:"trade_id"`
	Detail  []TradeDetail `json:"detail,failable"`
	// DetailList []TradeDetail `json:"detail"`
}

func main() {
	testcase1 := `{ "trade_id": "1234", "detail": { "item_id": "abcd" } }`
	testcase2 := `{ "trade_id": "1234", "detail": [{ "item_id": "abcd" }] }`

	var history TradeHistory
	if err := json.Unmarshal([]byte(testcase1), &history); err != nil {
		fmt.Println("testcase1 failed")
		fmt.Println(err)
	} else {
		fmt.Println("testcase1 passed")
	}

	fmt.Printf("history %#v\n", history)

	var history2 TradeHistory
	if err := json.Unmarshal([]byte(testcase2), &history2); err != nil {
		fmt.Println("testcase2 failed")
		fmt.Println(err)
	} else {
		fmt.Println("testcase2 passed")
	}
	fmt.Printf("history2 %#v\n", history2)
}
