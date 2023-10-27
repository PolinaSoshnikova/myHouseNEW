package createHouse

import (
	"fmt"
	"myHouse/infoHouse/infoappliances"
	"myHouse/infoHouse/infofamily"
	"myHouse/infoHouse/infofurniture"
	"myHouse/infoHouse/infosizes"
)

func CreateH() {
	sizes := infosizes.SizesIn()

	people := infofamily.PeopleIn()
	pets := infofamily.PetsIn()

	beds := infofurniture.BedsIn()
	chandeliers := infofurniture.ChandeliersIn()
	sofa := infofurniture.SofaIn()
	wardrobe := infofurniture.WardrobeIn()
	tables := infofurniture.TablesIn()

	cameras := infoappliances.CamerasIn()
	dishwasher := infoappliances.DishwasherIn()
	mixer := infoappliances.MixerIn()
	conditioners := infoappliances.ConditionersIn()
	blender := infoappliances.BlenderIn()

	fmt.Println("Мой дом")
	infosizes.SizesOut(sizes)

	fmt.Println("В нем живут:")
	infofamily.PetOut(pets)
	infofamily.PeopleOut(people)

	infofurniture.BedsOut(beds)
	infofurniture.ChandeliersOut(chandeliers)
	infofurniture.SofaOut(sofa)
	infofurniture.WardrobeOut(wardrobe)
	infofurniture.TablesOut(tables)

	infoappliances.CamerasOut(cameras)
	infoappliances.DishwasherOut(dishwasher)
	infoappliances.MixerOut(mixer)
	infoappliances.ConditionersOut(conditioners)
	infoappliances.BlenderOut(blender)
}
