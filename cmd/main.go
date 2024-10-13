package main

import (
	"goapi/controller"
	"goapi/db"
	"goapi/domain"
	"goapi/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if (err != nil) { panic(err) }

	ProductDb := infra.NewProductDb(dbConnection)

	Product := domain.NewProduct(ProductDb)

	productController := controller.NewProductController(Product)

	server.GET("/products", productController.GetProducts);
	server.GET("/product/:productId", productController.GetProductById);
	server.POST("/product", productController.CreateProduct);
	server.DELETE("/product/:productId", productController.DeleteProduct);
	server.PUT("/product/:productId", productController.UpdateProduct);

	server.Run(":8000")
}