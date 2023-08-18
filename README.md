# jsun

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.2.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

json encode/decode library

## Detail

jsun is a json decoding library created by extending encoding/json, a built-in module of Go.

Most go json libraries have limitations in unmarshal dynamic data.

```
{
    "trade_id": "1234",
    "detail": { "item_id": "abcd" }
}
or
{
    "trade_id": "1234",
    "detail": [{ "item_id": "abcd" }]
}
```

This type of json response is not desirable, but there are certainly cases.

## How to install

```
go get github.com/myyrakle/jsun@v0.2.0
```

## How to use

Usage is not that difficult.

Attach a failable tag to fields that may fail.

And for fields that can have polymorphic types, declare multiple Go struct fields and attach the same json tag name.

```
type TradeDetail struct {
	ItemID string `json:"item_id"`
}

type TradeHistory struct {
	TradeID    string        `json:"trade_id"`
	DetailList []TradeDetail `json:"detail,failable"`
	Detail     TradeDetail   `json:"detail,failable"`
}
```

Then, only fields that can be unmarshal will be assigned values, and the rest will be set to zero values.

The following code is a simple example.

```
package main

import (
	"fmt"

	json "github.com/myyrakle/jsun"
)

type TradeDetail struct {
	ItemID string `json:"item_id"`
}

type TradeHistory struct {
	TradeID    string        `json:"trade_id"`
	DetailList []TradeDetail `json:"detail,failable"`
	Detail     TradeDetail   `json:"detail,failable"`
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

```
