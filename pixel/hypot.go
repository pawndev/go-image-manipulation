package pixel

import "math"

func (p *Pixel) Hypot(p2 Pixel) Pixel {
	return Pixel{
		R: math.Hypot(p.R, p2.R),
		G: math.Hypot(p.G, p2.G),
		B: math.Hypot(p.B, p2.B),
		A: p.A,
	}
}

func (p *Pixel) Atan2(p2 Pixel) Pixel {
	return Pixel{
		R: math.Atan2(p.R, p2.R),
		G: math.Atan2(p.G, p2.G),
		B: math.Hypot(p.B, p2.B),
		A: p.A,
	}
}
