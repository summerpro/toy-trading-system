package types

type Amount int64

const ZeroAmount Amount = 0

func (a Amount) LessThan(b Amount) bool {
	return a < b
}

func (a Amount) LargerThan(b Amount) bool {
	return a > b
}

func (a Amount) EqualTo(b Amount) bool {
	return a == b
}

func (a Amount) LessThanOrEqual(b Amount) bool {
	return a <= b
}

func (a Amount) LargerThanOrEqual(b Amount) bool {
	return a >= b
}

func (a Amount) Add(b Amount) Amount {
	return a + b
}

func (a Amount) Sub(b Amount) Amount {
	return a - b
}

func (a Amount) Validate() bool {
	if a < 0 {
		return false
	}
	return true
}

func ToAmount(a int) Amount {
	return Amount(a)
}
