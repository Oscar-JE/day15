package vec

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDist(a Point, b Point) int {
	return abs(a.x-b.x) + abs(a.y+b.y)
}
