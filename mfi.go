package indicators

import "math"

// Money Flow Index
// https://www.investopedia.com/terms/m/mfi.asp
type MFI struct {
	n   float64
	i   float64
	tp  float64 // typical price
	pmf float64 // positive money flow
	nmf float64 // negative money flow
	mfi float64
}

func NewMFI(len int) *MFI {
	mfi := &MFI{float64(len), 0, math.NaN(), 0, 0, 50.0}
	return mfi
}

func (x *MFI) Update(o, h, l, c float64, v int64) {
	n := math.Max(x.n, x.i)
	tp := (h + l + c) / 3 // typical price
	mf := tp * float64(v) // money flow
	if x.tp < tp {
		x.pmf = (x.pmf*(n-1) + mf) / n
		x.nmf = x.nmf * (n - 1) / n
	} else {
		x.pmf = x.pmf * (n - 1) / n
		x.nmf = (x.nmf*(n-1) + mf) / n
	}
	if x.nmf == 0 {
		x.mfi = 100
	} else {
		mr := x.pmf / x.nmf // money ratio
		x.mfi = 100 - 100/(1+mr)
	}
	x.tp = tp
	x.i++
}

func (x *MFI) Value() float64 {
	return x.mfi
}
