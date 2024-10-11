package controller

import (
	"goapi/domain"
	"goapi/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	product domain.Product
}

func NewProductController(domain domain.Product) productController {
	return productController{
		product: domain,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.product.GetProducts()

	if (err != nil) {ctx.JSON(http.StatusInternalServerError, err)}

	ctx.JSON(http.StatusOK, products);
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.product.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct);
}