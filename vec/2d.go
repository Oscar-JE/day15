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

func (p Point) ManhattanDist(other Point) int {
	return abs(p.x-other.x) + abs(p.y+other.y)
}
