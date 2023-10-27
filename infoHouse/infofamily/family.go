package infofamily

import "fmt"

type People struct {
	name    string
	surname string
	age     float64
	height  float64
}

type Pets struct {
	name   string
	breed  string
	age    float64
	height float64
}

func PeopleIn() []People {
	woman1 := People{"Валерия", "Бондарь", 20, 170}
	woman2 := People{"Мария", "Афанасьева", 20, 162}
	woman3 := People{"Полина", "Сошникова", 20, 177}
	return []People{woman1, woman2, woman3}
}
func PetsIn() []Pets {
	cat := Pets{"Бадди", "шотландский", 5, 55}
	return []Pets{cat, cat}
}

func PeopleOut(person []People) {
	fmt.Println("Семья:")
	for _, people := range person {
		fmt.Print(people.name, " ", people.surname, " ", people.age, " ", people.height, "см\n")
	}
	fmt.Print("\n")
}
func PetOut(pets []Pets) {
	fmt.Println("Домашние животные:")
	for _, pet := range pets {
		fmt.Print(pet.name, " ", pet.breed, " ", pet.age, " ", pet.height, "см\n")
	}
	fmt.Print("\n")
}
