package img

import (
	"github.com/pawndev/golowpoly/pixel"
)

//const blur = *[9]float64{}

func (i *Image) MeanFilter() [][]pixel.Pixel {
	ps := make([][]pixel.Pixel, len(i.Pixels))
	for rowIndex := 0; rowIndex < i.Height; rowIndex++ {
		ps[rowIndex] = make([]pixel.Pixel, len(i.Pixels[rowIndex]))
		for colIndex := 0; colIndex < i.Width; colIndex++ {
			var sum int
			var totalDivide int = 0

			for rowindex2 := -1; rowindex2 <= 1; rowindex2++ {
				for colIndex2 := -1; colIndex2 <= 1; colIndex2++ {
					rowIndexSum := rowIndex + rowindex2
					if rowIndexSum >= 0 && rowIndexSum < len(i.Pixels) {
						colIndexSum := colIndex + colIndex2
						if colIndexSum >= 0 && colIndexSum < len(i.Pixels[rowIndexSum]) {
							totalDivide++
							sum = sum + i.Pixels[rowIndexSum][colIndexSum].G
						}
					}
				}
			}
			mean := sum / totalDivide
			ps[rowIndex][colIndex].Set("R", mean)
			ps[rowIndex][colIndex].Set("G", mean)
			ps[rowIndex][colIndex].Set("B", mean)
			ps[rowIndex][colIndex].Set("A", i.Pixels[rowIndex][colIndex].A)
		}
	}
	return ps
}
