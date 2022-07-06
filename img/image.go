package img

import (
	"errors"
	"github.com/pawndev/golowpoly/pixel"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var ErrImageTypeUnknown = errors.New("unknown image type")

type Image struct {
	Pixels [][]pixel.Pixel
	Width  int
	Height int
	rect   image.Rectangle
	raw    image.Image
}

func New(filePath string) (*Image, error) {
	s := strings.Split(filePath, ".")
	imgType := s[len(s)-1]

	switch imgType {
	case "jpeg", "jpg":
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	case "png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	default:
		return nil, errors.New("unknown image type")
	}

	imgReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(imgReader)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]pixel.Pixel
	for y := 0; y < height; y++ {
		var row []pixel.Pixel
		for x := 0; x < width; x++ {
			toPixel := pixel.RGBAToPixel(img.At(x, y).RGBA())
			row = append(row, toPixel)
		}
		pixels = append(pixels, row)
	}

	return &Image{
		Pixels: pixels,
		Width:  width,
		Height: height,
		rect:   img.Bounds(),
		raw:    img,
	}, nil
}

func (i *Image) Save(outputPath string) error {
	cimg := image.NewRGBA(i.rect)
	draw.Draw(cimg, i.rect, i.raw, image.Point{}, draw.Over)

	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			rowIndex, colIndex := y, x
			p := i.Pixels[rowIndex][colIndex]
			cimg.Set(x, y, color.RGBA{
				R: uint8(p.R),
				G: uint8(p.G),
				B: uint8(p.B),
				A: uint8(p.A),
			})
		}
	}

	s := strings.Split(outputPath, ".")
	imgType := s[len(s)-1]

	switch imgType {
	case "jpeg", "jpg", "png":
		fd, err := os.Create(outputPath)
		if err != nil {
			return err
		}

		switch imgType {
		case "jpeg", "jpg":
			if err = jpeg.Encode(fd, cimg, nil); err != nil {
				return err
			}
		case "png":
			if err = png.Encode(fd, cimg); err != nil {
				return err
			}
		}
	default:
		return ErrImageTypeUnknown
	}

	return nil
}
