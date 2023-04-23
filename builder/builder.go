package builder

import "fmt"

type IProductMaker interface {
	Reset()
	Design()
	ReDesign(string)
	BuyMaterial()
	DoSewage()
	PrintArtwork()
}

type ShirtMaker struct {
	shirt  Shirt
	design string
	price  int
}

func NewShirtMaker() *ShirtMaker {
	return &ShirtMaker{
		shirt:  Shirt{},
		design: "White flaky shirt",
		price:  100,
	}
}

func (s *ShirtMaker) Reset() {
	fmt.Printf("Reset Machine \n")
}

func (s *ShirtMaker) Design() {
	s.shirt.code = s.design
	fmt.Printf("The shirt is based on this design %s\n", s.design)
}

func (s *ShirtMaker) ReDesign(newDesign string) {
	s.design = newDesign
	s.shirt.code = s.design
	fmt.Printf("The shirt is based on this design %s\n", s.design)
}

func (s *ShirtMaker) BuyMaterial() {
	s.shirt.material = "cotton"
	fmt.Printf("The shirt is made of %s\n", s.shirt.material)
}

func (s *ShirtMaker) DoSewage() {
	s.shirt.name = s.design + " shirt"
	fmt.Printf("Process sewing...\n")
}

func (s *ShirtMaker) PrintArtwork() {
	s.shirt.price = float32(s.price) * 0.4
	fmt.Printf("On printing art...\n")
}

func (s *ShirtMaker) GetProduct() Shirt {
	return s.shirt
}

type TrouserMaker struct {
	trouser Trouser
	design  string
	price   int
}

func NewTrouserMaker() *TrouserMaker {
	return &TrouserMaker{
		trouser: Trouser{},
		design:  "Black long trousesr",
		price:   100,
	}
}

func (s *TrouserMaker) Reset() {
	fmt.Printf("Reset Machine\n")
}

func (s *TrouserMaker) Design() {
	s.trouser.code = s.design
	fmt.Printf("The trouser is based on this design %s\n", s.design)
}

func (s *TrouserMaker) ReDesign(newDesign string) {
	s.design = newDesign
	s.trouser.code = s.design
	fmt.Printf("The shirt is based on this design %s\n", s.design)
}

func (s *TrouserMaker) BuyMaterial() {
	s.trouser.material = "leather"
	fmt.Printf("The trouser is made of %s\n", s.trouser.material)
}

func (s *TrouserMaker) DoSewage() {
	s.trouser.name = s.trouser.material + " trousers"
	fmt.Printf("Process sewing...\n")
}

func (s *TrouserMaker) PrintArtwork() {
	s.trouser.price = float32(s.price) * 0.2
	fmt.Printf("On printing art...\n")
}

func (s *TrouserMaker) GetProduct() Trouser {
	return s.trouser
}
