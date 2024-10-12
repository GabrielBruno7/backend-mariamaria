package infra

import (
	"database/sql"
	"fmt"
	"goapi/model"
)

type ProductDb struct {
	connection *sql.DB
}

func NewProductDb(connection *sql.DB) ProductDb {
	return ProductDb{
		connection: connection,
	}
}

func (pr *ProductDb) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)

		if(err != nil) {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductDb) GetProductById(productId int ) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(productId).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if (err != nil) {
		if (err == sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &produto, nil
}

func (pr *ProductDb) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}