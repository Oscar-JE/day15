package parse

import (
	"bufio"
	"day15/sensorandbeacon"
	"day15/vec"
	"os"
	"strconv"
	"strings"
)

func Parse(filepath string) []sensorandbeacon.SensorAndBeacon {
	file, fileErr := os.Open(filepath)
	if fileErr != nil {
		panic("file open failed")
	}
	scanner := bufio.NewScanner(file)
	sensorPairs := []sensorandbeacon.SensorAndBeacon{}
	for scanner.Scan() {
		line := scanner.Text()
		sensor, beacon := parseLine(line)
		sensorPairs = append(sensorPairs, sensorandbeacon.Init(sensor, beacon))
	}
	return sensorPairs
}

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
