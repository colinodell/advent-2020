package utils

type Vector2 struct {
	X, Y int
}

type Vector3 struct {
	X, Y, Z int
}

type Vector4 struct {
	X, Y, Z, W int
}

type Vector interface {
	Nearby() []Vector
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

func (v *Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector3) Nearby() []Vector {
	ret, i := make([]Vector, 26), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x==0 && y==0 && z==0) {
					ret[i] = v.Add(Vector3{X: x, Y: y, Z: z})
					i++
				}
			}
		}
	}

	return ret
}

func (v *Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
		W: v.W + v2.W,
	}
}

func (v Vector4) Nearby() []Vector {
	ret, i := make([]Vector, 80), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						ret[i] = v.Add(Vector4{X: x, Y: y, Z: z, W: w})
						i++
					}
				}
			}
		}
	}

	return ret
}
