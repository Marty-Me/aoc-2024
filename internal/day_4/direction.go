package day4

type Direction struct {
	x int
	y int
}

var forward Direction = Direction{
	x: 1,
	y: 0,
}

var backward Direction = Direction{
	x: -1,
	y: 0,
}

var down Direction = Direction{
	x: 0,
	y: 1,
}

var up Direction = Direction{
	x: 0,
	y: -1,
}

var rightUp Direction = Direction{
	x: 1,
	y: -1,
}

var rightDown Direction = Direction{
	x: 1,
	y: 1,
}

var leftDown Direction = Direction{
	x: -1,
	y: 1,
}

var leftUp Direction = Direction{
	x: -1,
	y: -1,
}
