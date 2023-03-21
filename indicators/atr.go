package indicators

import (
	"math"
)

type TR struct {
	c   float64
	tr  float64
}

type MaxTR struct {
	tr  TR
	max float64
}

// Average True Range
// https://www.investopedia.com/terms/a/atr.asp
type ATR struct {
	n   float64
	i   float64
	tr  TR
	atr float64
}

// 14
func NewATR(len int) *ATR {
	atr := &ATR{float64(len), 0, TR{math.NaN(), math.NaN()}, math.NaN()}
	return atr
}

func NewMaxTR() *MaxTR {
	m := &MaxTR{TR{math.NaN(), math.NaN()}, 0.0}
	return m
}

func (x *TR) Update(o, h, l, c float64, v int64) {
	x.tr = h - l
	if !math.IsNaN(x.c) {
		tr1 := math.Abs(h - x.c)
		tr2 := math.Abs(l - x.c)
		x.tr = math.Max(x.tr, math.Max(tr1, tr2))
	}
	x.c = c
}

func (x *MaxTR) Update(o, h, l, c float64, v int64) {
	m := x.max
	x.tr.Update(o,h,l,c,v)
	x.max = math.Max(m, x.max)
}

func (x *ATR) Update(o, h, l, c float64, v int64) {
	x.tr.Update(o,h,l,c,v)

	// average true range calculation
	if math.IsNaN(x.atr) {
		x.atr = x.tr.tr
	} else {
		n := math.Min(x.i, x.n)
		x.atr = (x.atr*(n-1) + (x.tr.tr)) / n
	}
	x.i++
}

func (x *ATR) Value() float64 {
	return x.atr
}

func (x *MaxTR) Value() float64 {
	return x.max
}
