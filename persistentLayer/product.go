package persistentlayer


import "github.com/mohsenGoodarzi/webframework/entities"

func AllProducts() *models.Products {

	return &models.Products{
		Products: []models.Product{
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

func GetProduct(id int) *models.Product {
	return &models.Product{
		Id:   1,
		Name: "car",
	}
}
