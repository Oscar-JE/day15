package interval

type Interval struct {
	lowerBound int
	upperBound int
}

func Init(low int, heigh int) Interval {
	return Interval{low, heigh}
}

func (i Interval) Empty() bool {
	return i.NrOfElements() == 0
}

func (i Interval) NrOfElements() int { // open interval
	return relu(i.upperBound - i.lowerBound - 1)
}

func relu(x int) int {
	if x < 0 {
		return 0
	}
	return x
}

func (i Interval) IntersectsClosure(other Interval) bool {
	lower := i.InClosure(other.lowerBound)
	upper := i.InClosure(other.upperBound)
	lowerI := other.InClosure(i.lowerBound)
	upperI := other.InClosure(i.upperBound)
	return lower || upper || lowerI || upperI
}

func (i Interval) InClosure(element int) bool {
	return i.lowerBound <= element && element <= i.upperBound
}

func (i Interval) In(element int) bool {
	return i.lowerBound < element && element < i.upperBound
}

func (i *Interval) Absorb(other Interval) {
	i.lowerBound = min(i.lowerBound, other.lowerBound)
	i.upperBound = max(i.upperBound, other.upperBound)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
