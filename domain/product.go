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

func (pu *Product) GetProductById(productId int) (*model.Product, error) {
	product, err := pu.persistence.GetProductById(productId)

	if (err != nil) {
		return nil, err
	}

	return product, nil
}

func (pu *Product) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.persistence.CreateProduct(product)

	if (err != nil) {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *Product) DeleteProduct(productId int) (*model.Product, error) {
	product := pu.persistence.DeleteProduct(productId)

	return product, nil
}

func (pu *Product) UpdateProduct(productId int, productData *model.Product) (*model.Product, error) {
	updatedProduct, err := pu.persistence.UpdateProduct(productId, productData)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}