package util

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}

func (util *Util)GetMax(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
