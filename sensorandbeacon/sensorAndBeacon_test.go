package sensorandbeacon

import (
	"day15/interval"
	"day15/vec"
	"testing"
)

func TestIntervalAtRowTrivial(t *testing.T) { // är det här testet korrekt?
	sb := Init(vec.InitPoint(0, 0), vec.InitPoint(0, 0)) // ska nollan vara med här ?
	section := sb.ExcludedIntervalAtRow(0)
	if !section.Eq(interval.Init(-1, 1)) {
		t.Errorf("wtf")
	}
}

func TestIntervalAtRow(t *testing.T) {
	sb := Init(vec.InitPoint(0, 0), vec.InitPoint(0, 1))
	section := sb.ExcludedIntervalAtRow(0)
	if !section.Eq(interval.Init(-2, 2)) {
		t.Errorf("wrong interval calculated")
	}
}
