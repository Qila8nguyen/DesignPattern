package factory

import "fmt"

type StoreName string

const (
	HM     StoreName = "HM"
	Uniqlo StoreName = "Uniqlo"
)

type IFactory interface {
	CreateStore(storeName string) *IStore //factory method
}
type Factory struct{}

func (f *Factory) CreateStore(storeName string) (IStore, error) {
	if storeName == string(HM) {
		return NewHMStore(), nil
	}

	if storeName == string(Uniqlo) {
		return NewUniqloStore(), nil
	}

	return nil, fmt.Errorf("Brand is invalid")
}
