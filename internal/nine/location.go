package nine

type location struct {
	Value int
	Above *location
	Below *location
	Left  *location
	Right *location
}

func (l *location) Risk() int {
	return l.Value + 1
}

func (l *location) IsLowPoint() bool {
	isLow := true

	if l.Above != nil && l.Above.Value <= l.Value {
		isLow = false
	}
	if isLow && l.Below != nil && l.Below.Value <= l.Value {
		isLow = false
	}
	if isLow && l.Left != nil && l.Left.Value <= l.Value {
		isLow = false
	}
	if isLow && l.Right != nil && l.Right.Value <= l.Value {
		isLow = false
	}

	return isLow
}
