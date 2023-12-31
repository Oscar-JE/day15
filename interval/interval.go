package interval

import (
	"day15/integermath"
)

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

func (i Interval) Eq(other Interval) bool {
	return i.lowerBound == other.lowerBound && i.upperBound == other.upperBound
}

func (i Interval) NrOfElements() int { // open interval
	return integermath.Relu(i.upperBound - i.lowerBound - 1)
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

func (i Interval) Has(element int) bool {
	return i.lowerBound < element && element < i.upperBound
}

func (i *Interval) Absorb(other Interval) {
	i.lowerBound = min(i.lowerBound, other.lowerBound)
	i.upperBound = max(i.upperBound, other.upperBound)
}

func (i Interval) Split(element int) []Interval {
	if i.Has(element) {
		return []Interval{{i.lowerBound, element}, {element, i.upperBound}}
	}
	return []Interval{i}
}

func ShouldMerge(first Interval, second Interval) bool {
	lower := first.Has(second.lowerBound)
	upper := first.Has(second.upperBound)
	lowerI := second.Has(first.lowerBound)
	upperI := second.Has(first.upperBound)
	equals := first.Eq(second)
	return lower || upper || lowerI || upperI || equals
}

func (i *Interval) RemoveEdge(element int) {
	if element == i.lowerBound+1 {
		i.lowerBound += 1
	} else if element == i.upperBound-1 {
		i.upperBound -= 1
	}
}

func (i Interval) GetLowerBound() int {
	return i.lowerBound
}
func (i Interval) GetUpperBound() int {
	return i.upperBound
}
