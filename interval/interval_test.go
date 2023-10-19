package interval

import "testing"

func TestEmpty(t *testing.T) {
	e1 := Init(0, 0)
	if !e1.Empty() {
		t.Errorf("should be empty")
	}
	e2 := Init(0, 1)
	if !e2.Empty() {
		t.Errorf("should be empty")
	}
	e3 := Init(0, 2)

	if e3.Empty() {
		t.Errorf("should have one element")
	}
}

func TestIntersectsClosureContained(t *testing.T) {
	biggest := Init(0, 10)
	contained := Init(4, 7)
	if !biggest.IntersectsClosure(biggest) {
		t.Errorf("shoudl alwaus intersect own closure")
	}
	if !biggest.IntersectsClosure(contained) {
		t.Errorf("contained is in biggest")
	}
	if !contained.IntersectsClosure(biggest) {
		t.Errorf("should be a symetrik relation")
	}
}

func TestIntersectsClosureOverlapped(t *testing.T) {
	limitLow := Init(0, 5)
	limitHeigh := Init(5, 10)
	if !limitLow.IntersectsClosure(limitHeigh) {
		t.Errorf("a closure contains the sets boarder")
	}
	if !limitHeigh.IntersectsClosure(limitLow) {
		t.Errorf("symmetric")
	}
}
