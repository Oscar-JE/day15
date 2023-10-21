package sensorandbeacon

import (
	"day15/integermath"
	"day15/interval"
	"day15/vec"
	"fmt"
)

type SensorAndBeacon struct {
	sensor vec.Point
	beacon vec.Point
}

func Init(sensor vec.Point, beacon vec.Point) SensorAndBeacon {
	return SensorAndBeacon{sensor: sensor, beacon: beacon}
}

func (sb SensorAndBeacon) ExcludedIntervalAtRow(row int) interval.Interval {
	dist := vec.ManhattanDist(sb.sensor, sb.beacon)
	heightDif := integermath.Abs(sb.sensor.GetY() - row)
	halfLength := dist - heightDif + 1
	retInverval := interval.Init(sb.sensor.GetX()-halfLength, sb.sensor.GetX()+halfLength)
	/*
		 use if part one
		if sb.beacon.GetY() == row {
			retInverval.RemoveEdge(sb.beacon.GetX())
		}*/
	return retInverval
}

func (sb SensorAndBeacon) String() string {
	return fmt.Sprintf("s{%d,%d}-b{%d,%d}", sb.sensor.GetX(), sb.sensor.GetY(), sb.beacon.GetX(), sb.beacon.GetY())
}
