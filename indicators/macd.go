package indicators

import (
	"math"
)

type MACD struct {
	fma       EMA
	sma       EMA
	macd      float64
	signal    EMA
	histogram float64
}

func NewMACD(fast int, slow int, smoothing int) *MACD {
	result := &MACD{
		EMA{float64(fast), 0, math.NaN()},
		EMA{float64(slow), 0, math.NaN()},
		math.NaN(),
		EMA{float64(smoothing), 0, math.NaN()},
		math.NaN()}
	return result
}

func (x *MACD) Update(o, h, l, c float64, v int64) {
	x.fma.UpdateImpl(c)
	x.sma.UpdateImpl(c)
	x.macd = x.fma.ema - x.sma.ema
	x.signal.UpdateImpl(x.macd)
	x.histogram = x.macd - x.signal.ema
}

func (x *MACD) Value() (float64, float64, float64) {
	return x.macd, x.signal.Value(), x.histogram
}
