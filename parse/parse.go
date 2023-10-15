package parse

import (
	"day15/vec"
	"strconv"
	"strings"
)

func parseLine(line string) (vec.Point, vec.Point) {
	sensorAndBeacon := strings.Split(line, ":")
	sensor := sensorAndBeacon[0]
	beacon := sensorAndBeacon[1]
	return parseSensorOrBeacon(sensor), parseSensorOrBeacon(beacon)
}

func parseSensorOrBeacon(line string) vec.Point {
	infoStar := strings.Index(line, "at")
	information := line[infoStar:]
	xCoordAndYCoord := strings.Split(information, ",")
	xCoord := xCoordAndYCoord[0]
	x := extractInt(xCoord)
	yCoord := xCoordAndYCoord[1]
	y := extractInt(yCoord)
	return vec.InitPoint(x, y)
}

func extractInt(line string) int {
	x, error := strconv.Atoi(strings.Split(line, "=")[1])
	if error != nil {
		panic("critical error in parsing of file")
	}
	return x
}
