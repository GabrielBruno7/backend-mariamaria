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

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong", 
		})
	})

	server.GET("/products", productController.GetProducts);
	server.POST("/product", productController.CreateProduct);

	server.Run(":8000")
}