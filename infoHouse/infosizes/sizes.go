package infosizes

import "fmt"

type Sizes struct {
	name   string
	length float64
	width  float64
	height float64
}

func SizesIn() []Sizes {
	bedroom := Sizes{"спальня", 24, 10, 2.5}
	kitchen := Sizes{"кухня", 10, 10, 2.5}
	bathroom := Sizes{"ванная", 8, 8, 2.5}
	return []Sizes{bedroom, kitchen, bathroom}
}

func SizesOut(sizes []Sizes) {
	fmt.Print("Размеры комнаты(длина, ширина, высота):\n")
	for _, sizes := range sizes {
		fmt.Print(sizes.name, " ", sizes.length, sizes.width, sizes.height, "\n")
	}
}
