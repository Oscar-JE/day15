package vec

import "day15/integermath"

type Point struct {
	x int
	y int
}

func InitPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) Eq(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) GetX() int {
	return p.x
}

func (p Point) GetY() int {
	return p.y
}

func ManhattanDist(a Point, b Point) int {
	return integermath.Abs(a.x-b.x) + integermath.Abs(a.y-b.y)
}
