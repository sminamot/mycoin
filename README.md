# mycoin
## install
```
go get -u github.com/sminamot/mycoin
```

## sample
```go
package main

import (
	"fmt"
	"os"

	coinmarketcap "github.com/sminamot/coinmarketcap-go"
	linenotify "github.com/sminamot/line-notify-go"
	"github.com/sminamot/mycoin"
)

// Your portfolio
var cmap map[string]float64 = map[string]float64{
	"BTC": 1.23456,
	"ETH": 12.3456,
	// ...
}

func main() {
	co, err := coinmarketcap.GetAllCoinData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	coin := mycoin.NewCoins()

	for i, v := range cmap {
		coin.Add(i, v)
		coin.Set(i, co[i])
	}

	// Line Notify
	// see https://notify-bot.line.me/ja/
	n := linenotify.NewNotify()
	n.SetToken("<line notify token>")
	err = n.Notify(coin.MessageWithTime())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
```
