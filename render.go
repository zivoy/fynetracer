package main

import (
	"math"
	"renderShower/helpers"
	"sync"
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
	aspect := float64(width) / float64(height)
	var wg sync.WaitGroup
	for y := 0; y < height; y++ {
		wg.Add(1)
		go func(yLoc int) {
			defer wg.Done()

			yP := float64(yLoc) / float64(height)
			yP = -yP + 1

			yP = yP*2 - 1
			//row := buffer[yLoc]
			var xP float64
			for x := 0; x < width; x++ {
				xP = float64(x) * aspect / float64(width)
				xP = xP*2 - 1
				colour := calcPixel(xP, yP)

				buffer.SetRGBA(x, yLoc, colour.RGBA())
			}
			//time.Sleep(time.Second / 120)
		}(y)
	}
	wg.Wait()
	raster.Refresh()
}

var rayDirection = &helpers.Vector3{}

func calcPixel(x, y float64) helpers.Colour {
	rayDirection.X = x
	rayDirection.Y = y
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
	r := normal.Dot(lightDir.Mul(-1))
	col := sphereCol.MulCol(r)
	return col
}
