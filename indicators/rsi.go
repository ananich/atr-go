package indicators

import "math"

// Relative Strength Index
// https://www.investopedia.com/terms/r/rsi.asp
type RSI struct {
	n   float64
	i   float64
	cp  float64 // previous close
	ai  float64 // average increase
	ad  float64 // average decrease
	rsi float64
}

func NewRSI(len int) *RSI {
	rsi := &RSI{float64(len), 0, math.NaN(), 0, 0, 50.0}
	return rsi
}

func (x *RSI) Update(o, h, l, c float64, v int64) {
	if math.IsNaN(x.cp) {
		x.rsi = 0.5
	} else {
		n := math.Max(x.n, x.i)
		i := 0.0 // increase
		d := 0.0 // decrease
		if x.cp < c {
			i = c - x.cp
		} else {
			d = x.cp - c
		}
		x.ai = (x.ai*(n-1) + i) / n
		x.ad = (x.ad*(n-1) + d) / n
		if x.ad == 0 {
			x.rsi = 100
		} else {
			x.rsi = 100 - 100/(1+x.ai/x.ad)
		}
	}
	x.i++
}

func (x *RSI) Value() float64 {
	return x.rsi
}
