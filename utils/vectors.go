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
