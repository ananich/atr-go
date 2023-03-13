package indicators

import (
	"math"
)

type Indicator interface {
	Update(o, h, l, c float64, v int64)
	Value() float64
}

type ATR struct {
	n   float64
	i   float64
	c   float64
	atr float64
}

// Constructor
func NewATR(len int) *ATR {
	atr := &ATR{float64(len), 0, math.NaN(), math.NaN()}
	return atr
}

func (x *ATR) Update(o, h, l, c float64, v int64) {
	// true range calculation
	tr := h - l
	if !math.IsNaN(x.c) {
		tr1 := math.Abs(h - x.c)
		tr2 := math.Abs(l - x.c)
		tr = math.Max(tr, math.Max(tr1, tr2))
	}
	x.c = c

	// average true range calculation
	if math.IsNaN(x.atr) {
		x.atr = tr
	} else {
		n := math.Min(x.i, x.n)
		x.atr = (x.atr*(n-1) + (tr)) / n
	}
	x.i++
}

func (x *ATR) Value() float64 {
	return x.atr
}

type EMA struct {
	n   float64
	i   float64
	ema float64
}

const smoothing = 2.0

func NewEMA(len int) *EMA {
	ema := &EMA{float64(len), 0, math.NaN()}
	return ema
}

func (x *EMA) Update(o, h, l, c float64, v int64) {
	// exponential moving average calculation
	if math.IsNaN(x.ema) {
		x.ema = c
	} else {
		n := math.Min(x.i, x.n)
		k := smoothing / (n + 1)
		x.ema = x.ema*(1-k) + c*k
	}
	x.i++
}

func (x *EMA) Value() float64 {
	return x.ema
}
