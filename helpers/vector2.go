package helpers

import "math"

type Vector2 struct {
	X, Y float64
}

func NewVector2(x, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}
func NewVec2() *Vector2 {
	return NewVector2(0, 0)
}

func (v *Vector2) Normelize() *Vector2 {
	m := v.Magnitude()
	v.X /= m
	v.Y /= m
	return v
}
func (v *Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vector2) Dot(o *Vector2) float64 {
	return v.X*o.X + v.Y*o.Y
}
func (v *Vector2) Add(o *Vector2) *Vector2 {
	return NewVector2(v.X+o.X, v.Y+o.Y)
}
func (v *Vector2) Sub(o *Vector2) *Vector2 {
	return NewVector2(v.X-o.X, v.Y-o.Y)
}
func (v *Vector2) Mul(t float64) *Vector2 {
	return NewVector2(v.X*t, v.Y*t)
}
