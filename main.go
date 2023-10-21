package main

import (
	"day15/interval"
	"day15/parse"
	"day15/sensorandbeacon"
	"day15/vec"
	"fmt"
)

func main() {
	pairs := parse.Parse("input.txt")
	fmt.Println(pairs)
	// part one
	line := 10
	intervals := findIntervals(pairs, line)
	fmt.Println(nrOfExcludedAtRow(intervals, pairs, line))

	// part two
	fmt.Println(findTuningFrequency(pairs, 0, 4000000, 4000000))

	/*foundPoint
	fmt.Println(tuningFrequency(foundPoint, 20))*/
}

func findTuningFrequency(pairs []sensorandbeacon.SensorAndBeacon, min int, max int, tuningMultiple int) int {
	compareInterval := interval.Init(min, max)
	for i := min; i < max; i++ {
		intervals := findIntervals(pairs, i)
		/*fmt.Println(i)
		fmt.Println(intervals)*/
		for _, inter := range intervals {
			frequency := 0
			if compareInterval.Has(inter.GetLowerBound()) {
				frequency = tuningFrequency(vec.InitPoint(inter.GetLowerBound(), i), tuningMultiple)
			}
			if compareInterval.Has(inter.GetUpperBound()) {
				frequency = tuningFrequency(vec.InitPoint(inter.GetUpperBound(), i), tuningMultiple)
			}
			if frequency != 0 {

				return frequency
			}
		}

	}
	return 0
}

func findIntervals(sensorsAndBeacons []sensorandbeacon.SensorAndBeacon, row int) []interval.Interval {
	intervals := []interval.Interval{}
	for _, sensorAndBeacon := range sensorsAndBeacons {
		intervals = addToInterval(intervals, sensorAndBeacon.ExcludedIntervalAtRow(row))

	}
	return intervals
}

func nrOfExcludedAtRow(intervals []interval.Interval, sensorsAndBeacons []sensorandbeacon.SensorAndBeacon, row int) int {
	sum := 0
	for _, inter := range intervals {
		sum += inter.NrOfElements()
	}

	return sum
}

func addToInterval(intervals []interval.Interval, toAdd interval.Interval) []interval.Interval {
	retIntervals := []interval.Interval{}
	for _, inter := range intervals {
		if interval.ShouldMerge(toAdd, inter) {
			toAdd.Absorb(inter)
		} else {
			retIntervals = append(retIntervals, inter)
		}
	}
	retIntervals = append(retIntervals, toAdd)
	return retIntervals
}

func tuningFrequency(point vec.Point, tuningMultiple int) int {
	return point.GetX()*tuningMultiple + point.GetY()
}
