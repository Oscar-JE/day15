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

/*
 *               12 ########### 22
 * A   5---10
 * B   5-------------------20
 * C                  15---20
 * D   5-----------------------------------30
 * E                       20--------------30
 * F                                 25----30
 */

func TheStaticInterval() Interval {
	return Init(12, 22)
}

func TestShouldMergeNotA(t *testing.T) {
	GeneralTest(Init(5, 10), false, t)
}

func TestShouldMergeNotB(t *testing.T) {
	GeneralTest(Init(5, 20), true, t)
}

func TestShouldMergeNotC(t *testing.T) {
	GeneralTest(Init(15, 20), true, t)
}

func TestShouldMergeNotD(t *testing.T) {
	GeneralTest(Init(5, 30), true, t)
}

func TestShouldMergeNotE(t *testing.T) {
	GeneralTest(Init(20, 30), true, t)
}

func TestShouldMergeNotF(t *testing.T) {
	GeneralTest(Init(25, 30), false, t)
}

func TestShouldMergeNotG(t *testing.T) {
	GeneralTest(Init(5, 12), false, t)
}

func TestShouldMergeNotH(t *testing.T) {
	GeneralTest(Init(22, 30), false, t)
}

func TestShouldMergeNotI(t *testing.T) {
	GeneralTest(Init(5, 13), true, t)
}

func TestShouldMergeNotJ(t *testing.T) {
	GeneralTest(Init(21, 30), true, t)
}

func TestShouldMergeNotK(t *testing.T) {
	GeneralTest(Init(12, 22), true, t)
}

func GeneralTest(movingInterval Interval, expected bool, t *testing.T) {
	if ShouldMerge(TheStaticInterval(), movingInterval) != expected {
		t.Errorf("(%d, %d)", movingInterval.lowerBound, movingInterval.upperBound)
	}
}
