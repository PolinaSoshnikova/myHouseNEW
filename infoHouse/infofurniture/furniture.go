package infofurniture

import "fmt"

type Beds struct {
	name   string
	color  string
	length float64
	width  float64
}

type Chandeliers struct {
	name   string
	color  string
	length float64
	width  float64
}

type Sofa struct {
	name   string
	color  string
	length float64
	width  float64
}

type Wardrobe struct {
	name   string
	color  string
	length float64
	width  float64
}

type Tables struct {
	name   string
	color  string
	length float64
	width  float64
}

func BedsIn() []Beds {
	bed1 := Beds{"кровать Леры", "черный", 2, 1.5}
	bed2 := Beds{"кровать Маши", "розовый", 1.9, 1.5}
	bed3 := Beds{"кровать Полины", "радужный", 1.9, 1.5}
	return []Beds{bed1, bed2, bed3}
}
func ChandeliersIn() []Chandeliers {
	chandelier1 := Chandeliers{"компьютерное кресло", "серое", 0.5, 0.6}
	chandelier2 := Chandeliers{"стул", "синий", 0.5, 0.5}
	chandelier3 := Chandeliers{"стул", "красный", 0.5, 0.5}
	return []Chandeliers{chandelier1, chandelier2, chandelier3}
}
func SofaIn() []Sofa {
	sofa1 := Sofa{"диван", "красный", 2.5, 1.6}
	return []Sofa{sofa1}
}
func WardrobeIn() []Wardrobe {
	wardrobe1 := Wardrobe{"шкаф", "черный", 2.5, 1}
	return []Wardrobe{wardrobe1}
}
func TablesIn() []Tables {
	table1 := Tables{"стол", "коричневый", 1.3, 0.9}
	table2 := Tables{"маленький стол", "белый", 0.6, 0.4}
	return []Tables{table1, table2}

}

func BedsOut(beds []Beds) {
	fmt.Println("Кровати:")
	for _, oneOfBeds := range beds {
		fmt.Print(oneOfBeds.name, " ", oneOfBeds.color, " ", oneOfBeds.length, "м ", oneOfBeds.width, "м\n")
	}
	fmt.Print("\n")
}
func ChandeliersOut(chandeliers []Chandeliers) {
	fmt.Println("Люстры:")
	for _, oneOfChandeliers := range chandeliers {
		fmt.Print(oneOfChandeliers.name, " ", oneOfChandeliers.color, " ", oneOfChandeliers.length, "м ", oneOfChandeliers.width, "м\n")
	}
	fmt.Print("\n")
}
func SofaOut(sofas []Sofa) {
	fmt.Println("Диван:")
	for _, oneOfSofas := range sofas {
		fmt.Print(oneOfSofas.name, " ", oneOfSofas.color, " ", oneOfSofas.length, "м ", oneOfSofas.width, "м\n")
	}
	fmt.Print("\n")
}
func WardrobeOut(wardrobes []Wardrobe) {
	fmt.Println("Шкафы:")
	for _, oneOfWardrobes := range wardrobes {
		fmt.Print(oneOfWardrobes.name, " ", oneOfWardrobes.color, " ", oneOfWardrobes.length, "м ", oneOfWardrobes.width, "м\n")
	}
	fmt.Print("\n")
}
func TablesOut(tables []Tables) {
	fmt.Println("Столы:")
	for _, oneOfTables := range tables {
		fmt.Print(oneOfTables.name, " ", oneOfTables.color, " ", oneOfTables.length, "м ", oneOfTables.width, "м\n")
	}
	fmt.Print("\n")
}
