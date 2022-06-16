package service

import (
	"gautam/server/app/db/products_repo"
)

func GetProduct() ([]*products_repo.ProductsRetailers, error) {
	products, err := products_repo.GetProduct()
	if err != nil {
		return nil, err
	}
	return products, nil
}
