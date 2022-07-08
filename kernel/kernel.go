package kernel

type Kernel [][]float64

var BoxBlur Kernel = [][]float64{
	{1. / 9, 1. / 9, 1. / 9},
	{1. / 9, 1. / 9, 1. / 9},
	{1. / 9, 1. / 9, 1. / 9},
}

var GaussianBlur3x3 = [][]float64{
	{1. / 16, 2. / 16, 1. / 16},
	{2. / 16, 4. / 16, 2. / 16},
	{1. / 16, 2. / 16, 1. / 16},
}

var RidgeDetection Kernel = [][]float64{
	{-1., -1., -1.},
	{-1., -4., -1.},
	{-1., -1., -1.},
}

var Laplacien Kernel = [][]float64{
	{0., 1., 0.},
	{1., -4., 1.},
	{0., 1., 0.},
}
