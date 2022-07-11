package pixel

type Pixel struct {
	R float64
	G float64
	B float64
	A float64
}

func (p *Pixel) Set(keyName string, val float64) Pixel {
	switch keyName {
	case "R":
		p.R = val
	case "G":
		p.G = val
	case "B":
		p.B = val
	case "A":
		p.A = val
	}
	return *p
}

func RGBAToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		R: float64(r / 257),
		G: float64(g / 257),
		B: float64(b / 257),
		A: float64(a / 257),
	}
}
