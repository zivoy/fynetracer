package main

import (
	"math"
	"renderShower/helpers"
)

var (
	cameraPos = helpers.NewVector3(0, 0, -1)
	// cameraFacing = cameraPos.Sub(helpers.NewVec3()).Normelize()
	spherePos    = helpers.NewVector3(0, 0, 0)
	sphereRadius = .5
	sphereCol    = helpers.Colour{0, 1, 0, 1}
	lightDir     = helpers.NewVector3(1, 1, -1).Normelize()
)

func Render() {
	var colour helpers.Colour
	vec := &helpers.Vector2{}
	aspect := float64(width) / float64(height)
	for y := 0; y < height; y++ {
		vec.Y = float64(y) / float64(height)
		vec.Y = -vec.Y + 1

		vec.Y = vec.Y*2 - 1
		for x := 0; x < width; x++ {
			vec.X = float64(x) * aspect / float64(width)
			vec.X = vec.X*2 - 1

			colour = calcPixel(vec)
			buffer.SetRGBA(x, y, colour.RGBA())
		}
	}

	raster.Refresh()
}

var rayDirection = &helpers.Vector3{}

func calcPixel(pixel *helpers.Vector2) helpers.Colour {
	rayDirection.X = pixel.X
	rayDirection.Y = pixel.Y
	rayDirection.Z = -1

	a := rayDirection.Dot(rayDirection)
	b := 2 * cameraPos.Dot(rayDirection)
	c := cameraPos.Dot(cameraPos) - sphereRadius*sphereRadius
	d := b*b - 4*a*c
	if d < 0 {
		return helpers.Colour{0, 0, 0, 1}
	}
	sqt := math.Sqrt(d)
	dem := 2 * a
	t0 := (-b + sqt) / dem
	// t1 := (-b - sqt) / dem
	// tmin := math.Min(t0, t1)
	impact := cameraPos.Add(rayDirection.Mul(t0)) //(tmin)
	normal := impact.Sub(spherePos).Normelize()
	// col := helpers.ColourFromVec(normal.Mul(.5).Addf(.5))
	r := normal.Dot(lightDir.Mul(-1))
	col := sphereCol.MulCol(r)
	// return color.RGBA{uint8(normal.X * 255), uint8(normal.Y * 255), uint8(normal.Z * 255), 255}
	return col
}

// 	// for i := 0; i < 30000; i++ {
// 	// 	setRandomPixel()
// 	// }

// 	// circle -> x**2 + y**2 - r**2 = 0
// 	// line -> y = o_y + d_y * t
// 	// 		   x = o_x + d_x * t
// 	// (o_y + d_y * t)**2 + (o_x + d_x * t)**2 -r**2 = 0
// 	// o_y **2 + d_y**2 * t**2 + 2 * o_y * d_y * t + o_x **2 + d_x**2 * t**2 + 2 * o_x * d_x * t -r**2
// 	// (d_y**2 + d_x**2)*t**2 + (2*o_x*d_x + 2*o_y*d_y)*t + (o_y **2 +  o_x **2 -r**2)
// 	// t = (-b-+sqrt(b**2-4ac))/2a
// 	// d = (2*o_x*d_x + 2*o_y*d_y)**2 - 4 * ((d_y**2 + d_x**2) * (o_y **2 +  o_x **2 -r**2))
// 	// check val ^
// 	// t1 = (-(2*o_x*d_x + 2*o_y*d_y)+sqrt(d))/2* (d_y**2 + d_x**2)
// 	// t2 = (-(2*o_x*d_x + 2*o_y*d_y)-sqrt(d))/2* (d_y**2 + d_x**2)

// 	center := helpers.NewVector3(float64(width)/2, float64(height)/2, 0)

// 	c := color.RGBA{255, 0, 200, 255}
// 	p := center.Add(cameraPos)
// 	buffer.SetRGBA(int(p.X), int(p.Y), c)
// 	c.B = 150
// 	c.R = 30
// 	c.G = 200

// 	ea := cameraFacing.Y*cameraFacing.Y + cameraFacing.X*cameraFacing.X
// 	eb := 2*cameraPos.X*cameraFacing.X + 2*cameraPos.Y*cameraFacing.Y
// 	ec := cameraPos.Y*cameraPos.Y + cameraPos.X*cameraPos.X - sphereRadius*sphereRadius
// 	ed := eb*eb - 4*ea*ec
// 	if ed >= 0 {
// 		sqt := math.Sqrt(ed)
// 		dem := 2 * ea

// 		//t1
// 		t := (-eb + sqt) / dem
// 		// fmt.Print(t, " ")
// 		p = center.Add(cameraPos.Add(cameraFacing.Mul(t)))
// 		buffer.SetRGBA(int(p.X), int(p.Y), c)
// 		//t2
// 		t = (-eb - sqt) / dem
// 		// fmt.Println(t)
// 		p = center.Add(cameraPos.Add(cameraFacing.Mul(t)))
// 		buffer.SetRGBA(int(p.X), int(p.Y), c)
// 	}
// 	raster.Refresh()
// }

// func setRandomPixel() {
// 	buffer.SetRGBA(rand.Intn(width), rand.Intn(height), color.RGBA{uint8(rand.Intn(255)),
// 		uint8(rand.Intn(255)),
// 		uint8(rand.Intn(255)), 0xff})
// }
