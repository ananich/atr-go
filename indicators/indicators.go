package indicators

type Indicator interface {
	Update(o, h, l, c float64, v int64)
	Value() float64
}
