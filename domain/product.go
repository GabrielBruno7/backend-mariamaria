package domain

import (
	"goapi/infra"
	"goapi/model"
)

type Product struct {
	persistence infra.ProductDb
}

func NewProduct(persistence infra.ProductDb) Product {
	return Product{persistence: persistence}
}

func (pu *Product) GetProducts() ([]model.Product, error) {
	return pu.persistence.GetProducts()
}

func (pu *Product) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.persistence.CreateProduct(product)

	if (err != nil) {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}