package sensorandbeacon

import (
	"day15/interval"
	"day15/vec"
)

type SensorAndBeacon struct {
	sensor vec.Point
	beacon vec.Point
}

func Init(sensor vec.Point, beacon vec.Point) SensorAndBeacon {
	return SensorAndBeacon{sensor: sensor, beacon: beacon}
}

func ExcludedIntervalsAtRow(row int) []interval.Interval {

}
