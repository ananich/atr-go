package indicators

import (
	"math"
)

// Exponential Moving Average
// https://www.investopedia.com/ask/answers/122314/what-exponential-moving-average-ema-formula-and-how-ema-calculated.asp
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

func (x *EMA) calculate(c float64) float64 {
	if math.IsNaN(x.ema) {
		return c
	} else {
		n := math.Min(x.i, x.n)
		k := smoothing / (n + 1)
		return x.ema*(1-k) + c*k
	}
}

func (x *EMA) update(c float64) {
	x.ema = x.calculate(c)
	x.i++
}

func (x *EMA) Update(o, h, l, c float64, v int64) {
	x.update(c)
}

func (x *EMA) Value() float64 {
	return x.ema
}
