package main

import (
	"fmt"
	"github.com/pawndev/golowpoly/greyscale"
	"github.com/pawndev/golowpoly/img"
	"github.com/pawndev/golowpoly/kernel"
)

func main() {
	imgName := "building"
	extension := "jpg"
	i, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	i.Grayscale(greyscale.Luma)
	i.Pixels = i.Convolve(kernel.GaussianBlur3x3)
	i.Pixels = i.Convolve(kernel.Smoothing)
	ix := i.Convolve(kernel.SobelHorizontal)
	iy := i.Convolve(kernel.SobelVertical)
	//g := mathutils.Hypot(ix, iy)
	//a := mathutils.Atan2(ix, iy)

	if err := i.Save(fmt.Sprintf("./exports/%scanny-test.%s", imgName, extension)); err != nil {
		panic(err)
	}
}
