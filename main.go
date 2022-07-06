package main

import "github.com/pawndev/golowpoly/img"

func main() {
	i, _ := img.New("samples/cameraBruit.png")
	////i2 := i.Grayscale(greyscale.Luma)
	ps := i.MeanFilter()
	i.Pixels = ps
	if err := i.Save("./exports/cameraBruitMean.png"); err != nil {
		panic(err)
	}
	//fmt.Println(i.Height, i.Width)
	//fmt.Println(3 / 2)
}
