package main

import (
	"fmt"
	"projectname/main/builder"
	"projectname/main/factory"
	"projectname/main/observer"
)

func main() {
	FactoryVisualization()

	var singletonDirector *builder.Director
	BuilderVisualization(singletonDirector)

	ObserverVisualization()
}

func FactoryVisualization() {
	factory := factory.Factory{}
	HMStore, _ := factory.CreateStore("HM")
	Uniqlo, _ := factory.CreateStore("Uniqlo")

	fmt.Println(HMStore.GetName())
	Uniqlo.PrintLabel()
}

func BuilderVisualization(singletonDirector *builder.Director) {
	director := builder.GetInstance(singletonDirector)
	shirtMaker := builder.NewShirtMaker()
	director.MakeColorfulClothes(shirtMaker)

	var newShirt builder.Shirt
	newShirt = shirtMaker.GetProduct()
	println(newShirt.GetName())
	print("--------- \n")
	trouserMaker := builder.NewTrouserMaker()
	director.MakeSimpleClothes(trouserMaker)
	trouser := trouserMaker.GetProduct()

	println(trouser.GetName())

}

func ObserverVisualization() {
	shirtItem := observer.NewItem("Nike Shirt")

	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
