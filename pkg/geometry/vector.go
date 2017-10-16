package geometry

import "math"

type Vector interface {
	Length() float64
	GetComponents() []float64
}

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) GetComponents() []float64 {
	return []float64{v.X, v.Y}
}

type Vector3D struct {
	X, Y, Z float64
}

func (v Vector3D) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3D) GetComponents() []float64 {
	return []float64{v.X, v.Y, v.Z}
}
