package img

import (
	"github.com/pawndev/golowpoly/kernel"
	"github.com/pawndev/golowpoly/pixel"
)

//const blur = *[9]float64{}

func (i *Image) MeanFilter() [][]pixel.Pixel {
	p := make([][]pixel.Pixel, len(i.Pixels))
	for rowIndex := 0; rowIndex < i.Height; rowIndex++ {
		p[rowIndex] = make([]pixel.Pixel, len(i.Pixels[rowIndex]))
		for colIndex := 0; colIndex < i.Width; colIndex++ {
			var sumR int
			var sumG int
			var sumB int
			var totalDivide int = 0

			for rowindex2 := -1; rowindex2 <= 1; rowindex2++ {
				for colIndex2 := -1; colIndex2 <= 1; colIndex2++ {
					rowIndexSum := rowIndex + rowindex2
					if rowIndexSum >= 0 && rowIndexSum < len(i.Pixels) {
						colIndexSum := colIndex + colIndex2
						if colIndexSum >= 0 && colIndexSum < len(i.Pixels[rowIndexSum]) {
							totalDivide++
							sumR = sumR + i.Pixels[rowIndexSum][colIndexSum].R
							sumG = sumG + i.Pixels[rowIndexSum][colIndexSum].G
							sumB = sumB + i.Pixels[rowIndexSum][colIndexSum].B
						}
					}
				}
			}
			p[rowIndex][colIndex].Set("R", sumR/totalDivide)
			p[rowIndex][colIndex].Set("G", sumG/totalDivide)
			p[rowIndex][colIndex].Set("B", sumB/totalDivide)
			p[rowIndex][colIndex].Set("A", i.Pixels[rowIndex][colIndex].A)
		}
	}
	return p
}

func (i *Image) Convolute(k kernel.Kernel) [][]pixel.Pixel {
	p := make([][]pixel.Pixel, len(i.Pixels))
	for rowIndex := 0; rowIndex < i.Height; rowIndex++ {
		p[rowIndex] = make([]pixel.Pixel, len(i.Pixels[rowIndex]))
		for colIndex := 0; colIndex < i.Width; colIndex++ {
			var sumR float64
			var sumG float64
			var sumB float64
			kernelRowLen := (len(k) - 1) / 2

			for rowIndex2 := -kernelRowLen; rowIndex2 <= kernelRowLen; rowIndex2++ {
				kernelColLen := (len(k[kernelRowLen]) - 1) / 2
				for colIndex2 := -kernelColLen; colIndex2 <= kernelColLen; colIndex2++ {
					rowIndexSum := rowIndex + rowIndex2
					kernelRowIndex := rowIndex2 + kernelRowLen
					kernelColIndex := colIndex2 + kernelColLen
					if rowIndexSum >= 0 && rowIndexSum < len(i.Pixels) {
						colIndexSum := colIndex + colIndex2
						if colIndexSum >= 0 && colIndexSum < len(i.Pixels[rowIndexSum]) {
							sumR = sumR + (k[kernelRowIndex][kernelColIndex] * float64(i.Pixels[rowIndexSum][colIndexSum].R))
							sumG = sumG + (k[kernelRowIndex][kernelColIndex] * float64(i.Pixels[rowIndexSum][colIndexSum].G))
							sumB = sumB + (k[kernelRowIndex][kernelColIndex] * float64(i.Pixels[rowIndexSum][colIndexSum].B))
						}
					}
				}
			}
			p[rowIndex][colIndex].Set("R", int(sumR))
			p[rowIndex][colIndex].Set("G", int(sumG))
			p[rowIndex][colIndex].Set("B", int(sumB))
			p[rowIndex][colIndex].Set("A", i.Pixels[rowIndex][colIndex].A)
		}
	}
	return p
}
