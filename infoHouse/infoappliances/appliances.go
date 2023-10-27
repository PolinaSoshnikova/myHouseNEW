package infoappliances

import "fmt"

type Cameras struct {
	name   string
	color  string
	length float64
	width  float64
}

type Dishwasher struct {
	name   string
	color  string
	length float64
	width  float64
}

type Mixer struct {
	name   string
	color  string
	length float64
	width  float64
}

type Conditioners struct {
	name   string
	color  string
	length float64
	width  float64
}

type Blender struct {
	name   string
	color  string
	length float64
	width  float64
}

func CamerasIn() []Cameras {
	camera1 := Cameras{"камера", "синяя", 0.2, 0.1}
	camera2 := Cameras{"большая камера", "черная", 0.5, 0.4}
	return []Cameras{camera1, camera2}
}
func DishwasherIn() []Dishwasher {
	dishwasher1 := Dishwasher{"посудомойка", "белая", 0.3, 0.4}
	return []Dishwasher{dishwasher1}
}
func MixerIn() []Mixer {
	mixer1 := Mixer{"миксер", "белый", 0.2, 0.1}
	return []Mixer{mixer1}
}
func ConditionersIn() []Conditioners {
	conditioner1 := Conditioners{"кондиционер", "белый", 1, 0.4}
	conditioner2 := Conditioners{"маленький кондиционер", "серый", 0.6, 0.3}
	return []Conditioners{conditioner1, conditioner2}
}
func BlenderIn() []Blender {
	blender1 := Blender{"блендер", "черный", 0.3, 0.3}
	return []Blender{blender1}

}

func CamerasOut(cameras []Cameras) {
	fmt.Println("Камеры:")
	for _, oneOfCameras := range cameras {
		fmt.Print(oneOfCameras.name, " ", oneOfCameras.color, " ", oneOfCameras.length, "м ", oneOfCameras.width, "м\n")
	}
	fmt.Print("\n")
}
func DishwasherOut(dishwashers []Dishwasher) {
	fmt.Println("Посудомойка:")
	for _, oneOfDishwashers := range dishwashers {
		fmt.Print(oneOfDishwashers.name, " ", oneOfDishwashers.color, " ", oneOfDishwashers.length, "м ", oneOfDishwashers.width, "м\n")
	}
	fmt.Print("\n")
}
func MixerOut(mixers []Mixer) {
	fmt.Println("Миксер:")
	for _, oneOfMixers := range mixers {
		fmt.Print(oneOfMixers.name, " ", oneOfMixers.color, " ", oneOfMixers.length, "м ", oneOfMixers.width, "м\n")
	}
	fmt.Print("\n")
}
func ConditionersOut(conditioners []Conditioners) {
	fmt.Println("Кондиционеры:")
	for _, oneOfConditioners := range conditioners {
		fmt.Print(oneOfConditioners.name, " ", oneOfConditioners.color, " ", oneOfConditioners.length, "м ", oneOfConditioners.width, "м\n")
	}
	fmt.Print("\n")
}
func BlenderOut(blenders []Blender) {
	fmt.Println("Блендер:")
	for _, oneOfBlenders := range blenders {
		fmt.Print(oneOfBlenders.name, " ", oneOfBlenders.color, " ", oneOfBlenders.length, "м ", oneOfBlenders.width, "м\n")
	}
	fmt.Print("\n")
}
