package parse

import (
	"day15/vec"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
	sensor, beacon := parseLine(line)
	sExp := vec.InitPoint(2, 18)
	bExp := vec.InitPoint(-2, 15)
	if !sensor.Eq(sExp) {
		t.Errorf("Unexpected sensor parsing")
	}
	if !beacon.Eq(bExp) {
		t.Errorf("Unexpected breacon parsing")
	}
}
