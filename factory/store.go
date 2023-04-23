package factory

import "fmt"

type IStore interface {
	SetName(name string)
	GetName() string
	PrintLabel()
}

type HMStore struct {
	name string
}

type UniqloStore struct {
	name    string
	address string
}

func (s *HMStore) SetName(n string) {
	s.name = n
}

func (s *HMStore) GetName() string {
	return s.name
}

func (s *HMStore) PrintLabel() {
	fmt.Printf("Our %s name is %s", "store", s.name)
}

func (s *UniqloStore) SetName(n string) {
	s.name = n
}

func (s *UniqloStore) GetName() string {
	return s.name
}

func (s *UniqloStore) GetAddress() string {
	return s.address
}

func (s *UniqloStore) PrintLabel() {
	fmt.Printf("Our %s stores are placed accross the Asia", s.name)
}

func NewHMStore() IStore {
	return &HMStore{
		name: formatStoreName(string(HM)),
	}
}

func NewUniqloStore() IStore {
	return &UniqloStore{
		name:    formatStoreName(string(Uniqlo)),
		address: "Japan",
	}
}
