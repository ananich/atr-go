package indicators

import (
	"math"
)

// Directional Movement Index
// https://www.investopedia.com/terms/d/dmi.asp
// https://www.investopedia.com/terms/a/adx.asp
type DMI struct {
	n   float64
	i   float64
	hp  float64 // previous high
	lp  float64 // previous low
	atr ATR
	pDM float64
	nDM float64
	pDI float64 // positive directional indicator
	nDI float64 // negative directional indicator
	adx float64
}

func NewDMI(len int) *DMI {
	result := &DMI{
		float64(len),
		0,
		math.NaN(),
		math.NaN(),
		ATR{float64(len), 0, TR{math.NaN(), math.NaN()}, math.NaN()},
		0,
		0,
		math.NaN(),
		math.NaN(),
		0}
	return result
}

func (x *DMI) Update(o, h, l, c float64, v int64) {
	if !math.IsNaN(x.hp) {
		um := h - x.hp // up move
		dm := x.lp - l // down move
		pDM := 0.0
		if um > dm && um > 0 {
			pDM = um
		}
		mDM := 0.0
		if dm > um && dm > 0 {
			mDM = dm
		}
		n := math.Min(x.i, x.n)
		x.pDM = (x.pDM*(n-1) + pDM) / n
		x.nDM = (x.nDM*(n-1) + mDM) / n
		x.atr.Update(o, h, l, c, v)
		x.pDI = x.pDM * 100 / x.atr.Value()
		x.nDI = x.nDM * 100 / x.atr.Value()
		dx := 100 * math.Abs(x.pDI-x.nDI) / math.Abs(x.pDI+x.nDI)
		x.adx = (x.adx*(n-1) + dx) / n
	}
	x.i++
	x.hp = h
	x.lp = l
}

func (x *DMI) Value() (float64, float64, float64) {
	return x.pDI, x.nDI, x.adx
}
