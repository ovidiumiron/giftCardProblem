package priceProduct

import "fmt"

type PriceProduct struct {
	price   int
	product string
}

func NewPriceProduct(price int, product string) (p PriceProduct, err error) {
	if price < 0 {
		err = InvalidPriceError(price)
		return
	}
	return PriceProduct{price, product}, nil
}

func (p PriceProduct) Price() int {
	return p.price
}

func (p PriceProduct) Product() string {
	return p.product
}

func (p PriceProduct) String() string {
	return fmt.Sprintf("%d %s", p.price, p.product)
}

type InvalidPriceError int

func (i InvalidPriceError) Error() string {
	return fmt.Sprintf("The price %q is invalid because is less then 0.", i)
}
