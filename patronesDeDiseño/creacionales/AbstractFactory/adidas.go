// Fábrica concreta
// Cada fábrica concreta (Adidas y Nike) implementa los métodos para crear productos específicos de su marca.
package main

type Adidas struct {
}

/*Se esta creando un Shoe de tipo adidas y en func (a *Adidas) makeShoe() IShoe, se esta obligando a que ese objeto, si
o si implemente los metdos de la inferfaz IShoe, en este caso, son solo getters y setters
*/
func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
