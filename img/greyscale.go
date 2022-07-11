package img

import (
	"github.com/pawndev/golowpoly/greyscale"
	"sync"
)

// Grayscale turns the images to grayscale.
func (i *Image) Grayscale(algorithm greyscale.Algorithm) *Image {
	var wg sync.WaitGroup

	for rowIndex := 0; rowIndex < i.Height; rowIndex++ {
		wg.Add(1)

		go (func(rowIndex int, img *Image) {
			for colIndex := 0; colIndex < img.Width; colIndex++ {
				pixel := img.Pixels[rowIndex][colIndex]

				var gray float64
				if algorithm == greyscale.Luma {
					gray = float64(pixel.R)*0.2126 + float64(pixel.G)*0.7152 + float64(pixel.B)*0.0722
				} else if algorithm == greyscale.Desaturation {
					gray = (max(pixel.R, pixel.G, pixel.B) + min(pixel.R, pixel.G, pixel.B)) / 2
				} else {
					gray = (pixel.R + pixel.G + pixel.B) / 3
				}
				pixel.Set("R", gray)
				pixel.Set("G", gray)
				pixel.Set("B", gray)

				img.Pixels[rowIndex][colIndex] = pixel
			}
			wg.Done()
		})(rowIndex, i)
	}

	wg.Wait()
	return i
}

// max finds max among multiple integers
func max(ints ...float64) float64 {
	var max = ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

// min finds min among multiple integers
func min(ints ...float64) float64 {
	var min = ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}
	return min
}
