package models

type Product struct {
	Id   int
	Name string
}
type Products struct {
	Products []Product
}

func AllProducts() *Products {

	return &Products{
		Products: []Product{
			{
				Id:   1,
				Name: "car",
			},
			{
				Id:   2,
				Name: "bike",
			},
		},
	}
}

func GetProduct(id int) *Product {
	return &Product{
		Id:   1,
		Name: "car",
	}
}
