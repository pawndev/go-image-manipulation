package greyscale

type Algorithm int

const (
	Average Algorithm = iota
	Luma
	Desaturation
)
