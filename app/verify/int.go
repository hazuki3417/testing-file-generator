package verify

type IntVerify struct {
	Value int
}

func (iv *IntVerify) Negative() bool {
	return iv.Value < 0
}

func (iv *IntVerify) Positive() bool {
	return 0 < iv.Value
}

func (iv *IntVerify) Equal(size int) bool {
	return size == iv.Value
}

func (iv *IntVerify) NotEqual(size int) bool {
	return size != iv.Value
}

func (iv *IntVerify) Gt(size int) bool {
	return size < iv.Value
}

func (iv *IntVerify) Ge(size int) bool {
	return size <= iv.Value
}

func (iv *IntVerify) Le(size int) bool {
	return iv.Value <= size
}

func (iv *IntVerify) Lt(size int) bool {
	return iv.Value < size
}

func (sv *IntVerify) InRange(min int, max int) bool {
	return min <= sv.Value && sv.Value <= max
}

func (sv *IntVerify) OutRange(min int, max int) bool {
	return sv.Value < min || max < sv.Value
}
