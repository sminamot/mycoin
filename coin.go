package mycoin

import (
	"fmt"
	"strings"
	"time"

	coinmarketcap "github.com/sminamot/coinmarketcap-go"
)

const NOTIFY_TIME_FORMAT = "15:04"

type MyCoins map[string]*MyCoin

type MyCoin struct {
	Total      float64
	Yen        float64
	Rate       float64
	Changed24h float64
}

func NewCoins() *MyCoins {
	return &MyCoins{}
}

func (m MyCoins) Add(s string, total float64) {
	m[s] = &MyCoin{
		Total: total,
	}
}

func (m MyCoins) Set(s string, b coinmarketcap.Coin) {
	if my, ok := m[b.Symbol]; ok {
		my.Yen = my.Total * b.PriceJpy
		my.Rate = b.PriceJpy
		my.Changed24h = b.PercentChange24H
	}
}

func (m MyCoins) SetYen(s string, y float64) {
	if my, ok := m[s]; ok {
		my.Yen = y
	}
}

func (m MyCoins) SetChanged24h(s string, c float64) {
	if my, ok := m[s]; ok {
		my.Changed24h = c
	}
}

func (m MyCoins) SetRate(s string, r float64) {
	if my, ok := m[s]; ok {
		my.Rate = r
	}
}

func (m MyCoins) TotalYen() (t float64) {
	for _, v := range m {
		t += v.Yen
	}
	return t
}

func (m MyCoins) Message() string {
	return m.message(false)
}

func (m MyCoins) MessageWithTime() string {
	return m.message(true)
}

func (m MyCoins) message(wt bool) string {
	ss := make([]string, 0, len(m)*4+4)
	if wt {
		ss = append(ss, fmt.Sprint(time.Now().Format(NOTIFY_TIME_FORMAT)))
	}
	totalYen := m.TotalYen()
	for i, v := range m {
		ss = append(ss, fmt.Sprint("■"+i))
		ss = append(ss, fmt.Sprintf("total: %.5f", v.Total))
		ss = append(ss, fmt.Sprint("+-(%): ", v.Changed24h))
		ss = append(ss, fmt.Sprintf("rate: %.5f", v.Rate))
		ss = append(ss, fmt.Sprintf("yen: %.5f", v.Yen))
	}
	ss = append(ss, fmt.Sprint("■total"))
	ss = append(ss, fmt.Sprintf("yen: %.5f", totalYen))
	return strings.Join(ss, "\n")
}
