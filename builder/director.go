package builder

import "sync"

type Director struct{}

var mu sync.Mutex

func (d *Director) MakeSimpleClothes(builder IProductMaker) {
	builder.Reset()
	builder.Design()
	builder.BuyMaterial()
	builder.DoSewage()
}

func (d *Director) MakeColorfulClothes(builder IProductMaker) {
	builder.Reset()
	builder.Design()
	builder.ReDesign("modern")
	builder.BuyMaterial()
	builder.DoSewage()
	builder.PrintArtwork()
}

func GetInstance(instance *Director) *Director {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Director{}
		return instance
	} else {
		return instance
	}
}
