package indicators

import (
	"math"
)

// Average True Range
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
