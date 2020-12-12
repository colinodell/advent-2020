package utils

type Vector2 struct {
	X, Y int
}

func (v *Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector2) Multiply(i int) Vector2 {
	return Vector2{
		X: v.X * i,
		Y: v.Y * i,
	}
}

func (v Vector2) Reduce() Vector2 {
	gcd := GCD(v.X, v.Y)

	return Vector2{
		X: v.X / gcd,
		Y: v.Y / gcd,
	}
}

func (v Vector2) ManhattanDistance() int {
	return Abs(v.X) + Abs(v.Y)
}

func (v Vector2) RotateClockwise(quarterTurns int) Vector2 {
	for quarterTurns > 0 {
		v.X, v.Y = v.Y, v.X
		v.Y = -v.Y
		quarterTurns--
	}

	return v
}

func (v Vector2) RotateCounterClockwise(quarterTurns int) Vector2 {
	for quarterTurns > 0 {
		v.X, v.Y = v.Y, v.X
		v.X = -v.X
		quarterTurns--
	}

	return v
}
