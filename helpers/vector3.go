package helpers

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}
func NewVec3() *Vector3 {
	return NewVector3(0, 0, 0)
}

func (v *Vector3) Normelize() *Vector3 {
	m := v.Magnitude()
	v.X /= m
	v.Y /= m
	v.Z /= m
	return v
}
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
func (v *Vector3) Dot(o *Vector3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}
func (v *Vector3) Cross(o *Vector3) *Vector3 {
	x := v.Y*o.Z - v.Z*o.Y
	y := v.Z*o.X - v.X*o.Z
	z := v.X*o.Y - v.Y*o.X
	return NewVector3(x, y, z)
}
func (v *Vector3) Add(o *Vector3) *Vector3 {
	return NewVector3(v.X+o.X, v.Y+o.Y, v.Z+o.Z)
}
func (v *Vector3) Sub(o *Vector3) *Vector3 {
	return NewVector3(v.X-o.X, v.Y-o.Y, v.Z-o.Z)
}
func (v *Vector3) Addf(f float64) *Vector3 {
	return NewVector3(v.X+f, v.Y+f, v.Z+f)
}
func (v *Vector3) Subf(f float64) *Vector3 {
	return NewVector3(v.X-f, v.Y-f, v.Z-f)
}
func (v *Vector3) Mul(t float64) *Vector3 {
	return NewVector3(v.X*t, v.Y*t, v.Z*t)
}
