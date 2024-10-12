package controller

import (
	"goapi/domain"
	"goapi/model"
	"strconv"

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

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if (id == "") {
		response := model.Response{Message: "Product Id can't be null"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if (err!=nil) {
		response := model.Response{Message: "Product Id should be a number"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.product.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{Message: "Product not found in the database"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product);
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