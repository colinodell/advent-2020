package utils

type Vector2 struct {
	X, Y int
}

func (v Vector2) Reduce() Vector2 {
	gcd := GCD(v.X, v.Y)

	return Vector2{
		X: v.X / gcd,
		Y: v.Y / gcd,
	}
}
