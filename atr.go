package indicators

import (
	"math"
)

// --- True Range --- //

type TR struct {
	c  float64
	tr float64
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

func (x *TR) Value() float64 {
	return x.tr
}

// --- Max True Range --- //

type MaxTR struct {
	tr  TR
	max float64
}

func NewMaxTR() *MaxTR {
	m := &MaxTR{TR{math.NaN(), math.NaN()}, 0.0}
	return m
}

func (x *MaxTR) Update(o, h, l, c float64, v int64) {
	x.tr.Update(o, h, l, c, v)
	x.max = math.Max(x.max, x.tr.tr)
}

func (x *MaxTR) Value() float64 {
	return x.max
}

//  --- Average True Range --- //
// https://www.investopedia.com/terms/a/atr.asp

type ATR struct {
	n   float64
	i   float64
	tr  TR
	atr float64
}

// len is 14 by default
func NewATR(len int) *ATR {
	atr := &ATR{float64(len), 0, TR{math.NaN(), math.NaN()}, math.NaN()}
	return atr
}

func (x *ATR) Update(o, h, l, c float64, v int64) {
	x.tr.Update(o, h, l, c, v)
	if math.IsNaN(x.atr) {
		x.atr = x.tr.Value()
	} else {
		n := math.Min(x.i, x.n)
		x.atr = (x.atr*(n-1) + (x.tr.Value())) / n
	}
	x.i++
}

func (x *ATR) Value() float64 {
	return x.atr
}

// --- Average True Range Percent --- //

type ATRP struct {
	atr  ATR
	atrp float64
}

// len is 14 by default
func NewATRP(len int) *ATRP {
	atrp := &ATRP{ATR{float64(len), 0, TR{math.NaN(), math.NaN()}, math.NaN()}, math.NaN()}
	return atrp
}

func (x *ATRP) Update(o, h, l, c float64, v int64) {
	x.atr.Update(o, h, l, c, v)
	x.atrp = x.atr.Value() / c
}

func (x *ATRP) Value() float64 {
	return x.atrp
}
