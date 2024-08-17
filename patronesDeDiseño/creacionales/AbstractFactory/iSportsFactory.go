// Interfaz de la fábrica abstracta
package main

import "fmt"

//ISportsFactory: Es la interfaz de fábrica abstracta que define métodos para crear productos relacionados, en este caso makeShoe y makeShirt.
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
