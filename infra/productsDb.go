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
	query := "SELECT id, name, price FROM product ORDER BY id"

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

func (pr *ProductDb) DeleteProduct(productId int) (*model.Product) {
    query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")

    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer query.Close()

    _, err = query.Exec(productId)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    return &model.Product{}
}


func (pr *ProductDb) UpdateProduct(productId int, productData *model.Product) (*model.Product, error) {
	query := "UPDATE product SET name = $2, price = $3 WHERE id = $1 RETURNING id, name, price"

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(productId, productData.Name, productData.Price)

	var updatedProduct model.Product
	if err := row.Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return &updatedProduct, nil
}
