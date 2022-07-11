package mathutils

import "github.com/pawndev/golowpoly/pixel"

func Hypot(m [][]pixel.Pixel, m2 [][]pixel.Pixel) [][]pixel.Pixel {
	p := make([][]pixel.Pixel, len(m))
	for i, _ := range m {
		p[i] = make([]pixel.Pixel, len(m[i]))
		for j, localPixel := range m[i] {
			p[i][j] = localPixel.Hypot(m2[i][j])
		}
	}

	return p
}

func Atan2(m [][]pixel.Pixel, m2 [][]pixel.Pixel) [][]pixel.Pixel {
	p := make([][]pixel.Pixel, len(m))
	for i, _ := range m {
		p[i] = make([]pixel.Pixel, len(m[i]))
		for j, localPixel := range m[i] {
			p[i][j] = localPixel.Atan2(m2[i][j])
		}
	}

	return p
}
