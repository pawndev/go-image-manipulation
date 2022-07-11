package main

import (
	"fmt"
	"github.com/pawndev/golowpoly/greyscale"
	"github.com/pawndev/golowpoly/img"
	"github.com/pawndev/golowpoly/kernel"
)

func main() {
	imgName := "sample1"
	extension := "png"
	i, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ps := i.MeanFilter()
	i.Pixels = ps
	if err := i.Save(fmt.Sprintf("./exports/%sMean.%s", imgName, extension)); err != nil {
		panic(err)
	}

	i1, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ps1 := i1.Convolve(kernel.BoxBlur)
	i1.Pixels = ps1
	if err := i1.Save(fmt.Sprintf("./exports/%sBoxBlur.%s", imgName, extension)); err != nil {
		panic(err)
	}

	i2, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ps2 := i2.Convolve(kernel.GaussianBlur3x3)
	i2.Pixels = ps2
	if err := i2.Save(fmt.Sprintf("./exports/%sConvolGaussian3x3.%s", imgName, extension)); err != nil {
		panic(err)
	}

	i3, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ps3 := i3.Convolve(kernel.RidgeDetection)
	i3.Pixels = ps3
	if err := i3.Save(fmt.Sprintf("./exports/%sConvolRidgeDetection.%s", imgName, extension)); err != nil {
		panic(err)
	}

	ij, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ij = ij.Grayscale(greyscale.Luma)
	psj := ij.Convolve(kernel.RidgeDetection)
	ij.Pixels = psj
	if err := ij.Save(fmt.Sprintf("./exports/%sRidgeDetectionAfterGreyscale.%s", imgName, extension)); err != nil {
		panic(err)
	}

	ik, _ := img.New(fmt.Sprintf("samples/%s.%s", imgName, extension))
	ik = ij.Grayscale(greyscale.Luma)
	psk := ik.Convolve(kernel.Laplacien)
	ik.Pixels = psk
	if err := ik.Save(fmt.Sprintf("./exports/%sLaplacienAfterGreyscale.%s", imgName, extension)); err != nil {
		panic(err)
	}
}
