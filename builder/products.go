package builder

type IProduct interface {
	GetPrice() float32
	GetName() string
}

type Shirt struct {
	price    float32
	code     string
	name     string
	material string
}

func (s *Shirt) GetPrice() float32 {
	return s.price
}

func (s *Shirt) GetName() string {
	return s.name
}

type Trouser struct {
	price       float32
	code        string
	name        string
	trouserType string
	material    string
}

func (s *Trouser) GetPrice() float32 {
	return s.price
}

func (s *Trouser) GetName() string {
	return s.name
}
