package solid

import "fmt"

type IClothes interface {
	GetPrice() float32
}

func NewShirt() *Shirt {
	return &Shirt{
		price:    30.23,
		code:     "S32-w3",
		name:     "White shirt",
		material: "cotton",
	}
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

type Invoice struct {
	clothes    IClothes
	quantity   int
	taxRate    float32
	totalPrice float32
}

func NewInvoice(clothes IClothes) *Invoice {
	return &Invoice{
		clothes:    clothes,
		quantity:   5,
		totalPrice: 0,
		taxRate:    0.1,
	}
}
func (i *Invoice) CalculateTotal() float32 {
	i.totalPrice = float32(i.quantity) * i.clothes.GetPrice() * (1 + i.taxRate)
	return i.totalPrice
}

// @violates -> what if we change the printing format?
// func (i *Invoice) PrintInvoice() {
// 	fmt.Printf("Tax rate is %f\n", i.taxRate)
// 	fmt.Printf("Total quantity is %d\n", i.quantity)
// }
// -----------

type InvoicePrinter struct {
	invoice *Invoice
}

func (ip *InvoicePrinter) AttachInvoice(invoice *Invoice) {
	ip.invoice = invoice
}

func (i *InvoicePrinter) PrintInvoice() {
	fmt.Printf("Tax rate is %f\n", i.invoice.taxRate)
	fmt.Printf("Total quantity is %d\n", i.invoice.quantity)
}

func mainSingle() {
	shirt := NewShirt()
	invoice := NewInvoice(shirt)
	invoicePrinter := &InvoicePrinter{
		invoice: invoice,
	}

	invoicePrinter.PrintInvoice()
}
