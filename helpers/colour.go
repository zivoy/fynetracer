package helpers

import (
	"image/color"
)

type Colour struct {
	R, G, B, A float64
}

func (c Colour) RGBA() color.RGBA {
	r := clamp(c.R) * 0xff
	g := clamp(c.G) * 0xff
	b := clamp(c.B) * 0xff
	a := clamp(c.A) * 0xff
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func (c Colour) MulCol(f float64) Colour {
	f = clamp(f)
	return Colour{c.R * f, c.G * f, c.B * f, c.A}
}

func ColourFromVec(v *Vector3) Colour {
	return Colour{v.X, v.Y, v.Z, 1}
}
func ColourFromRGBA(c color.RGBA) Colour {
	return Colour{float64(c.R) / 0xff, float64(c.G) / 0xff, float64(c.B) / 0xff, float64(c.A) / 0xff}
}
func ColourFromColor(c color.Color) Colour {
	r, g, b, a := c.RGBA()
	return Colour{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff, float64(a) / 0xffff}
}

func clamp(f float64) float64 {
	if f < 0 {
		return 0
	}
	if f > 1 {
		return 1
	}
	return f
	//return math.Min(math.Max(f, 0), 1)
}
