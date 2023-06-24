package indicators

import (
	"math"
)

// Rate of Change
// https://www.investopedia.com/terms/p/pricerateofchange.asp
type RoC struct {
	src []float64
	i   int
	roc float64
}

func NewRoC(len int) *RoC {
	roc := &RoC{make([]float64, len), 0, math.NaN()}
	for i := 0; i < len; i++ {
		roc.src[i] = math.NaN()
	}
	return roc
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (x *RoC) update(c float64) {
	n := len(x.src) - 1 - min(x.i, len(x.src)-1)
	s := x.src[n]
	x.roc = (c - s) / s * 100
	x.src = x.src[1:]
	x.src = append(x.src, c)
	x.i++
}

func (x *RoC) Update(o, h, l, c float64, v int64) {
	x.update(c)
}

func (x *RoC) Value() float64 {
	return x.roc
}
