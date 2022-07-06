package pixel

type Pixel struct {
	R int
	G int
	B int
	A int
}

func (p *Pixel) Set(keyName string, val int) Pixel {
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
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257),
	}
}
