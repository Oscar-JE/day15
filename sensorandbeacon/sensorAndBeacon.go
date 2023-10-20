package sensorandbeacon

import (
	"day15/integermath"
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

func (sb SensorAndBeacon) ExcludedIntervalsAtRow(row int) interval.Interval {
	dist := vec.ManhattanDist(sb.sensor, sb.beacon)
	heightDif := integermath.Abs(sb.sensor.GetY())
	halfLength := dist - heightDif + 1
	return interval.Init(sb.sensor.GetX()-halfLength, sb.sensor.GetX()+halfLength)
}
