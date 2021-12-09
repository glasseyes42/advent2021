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

func checkIfContains(list []*location, l *location) bool {
	for _, existing := range list {
		if existing == l {
			return true
		}
	}
	return false
}

func (l *location) DiscoverBasinEdges(soFar []*location) []*location {
	list := make([]*location, len(soFar))
	copy(list, soFar)

	if l.Above != nil && l.Above.Value < 9 && !checkIfContains(list, l.Above) {
		list = append(list, l.Above)
		list = l.Above.DiscoverBasinEdges(list)
	}
	if l.Below != nil && l.Below.Value < 9 && !checkIfContains(list, l.Below) {
		list = append(list, l.Below)
		list = l.Below.DiscoverBasinEdges(list)
	}
	if l.Left != nil && l.Left.Value < 9 && !checkIfContains(list, l.Left) {
		list = append(list, l.Left)
		list = l.Left.DiscoverBasinEdges(list)
	}
	if l.Right != nil && l.Right.Value < 9 && !checkIfContains(list, l.Right) {
		list = append(list, l.Right)
		list = l.Right.DiscoverBasinEdges(list)
	}

	return list
}

func (l *location) BasinSize() int {
	list := l.DiscoverBasinEdges([]*location{l})

	return len(list)
}
