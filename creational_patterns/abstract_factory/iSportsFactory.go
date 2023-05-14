package abstract_factory

import "fmt"

type ISSportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportFactory(brand string) (ISSportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
