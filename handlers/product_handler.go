package handlers

import (
	"backend1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ALL PRODUCTS
func GetAllProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "List of products",
		Result:  models.ListProduct,
	})
}

// GET PRODUCT BY ID
func GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, product := range models.ListProduct {
		if product.ID == id {
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Product found",
				Result:  product,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "Product not found",
	})
}

// CREATE PRODUCT
func CreateProduct(ctx *gin.Context) {
	var data models.Product

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	models.ProductCounter++
	data.ID = strconv.Itoa(models.ProductCounter)

	models.ListProduct = append(models.ListProduct, data)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Product created successfully",
		Result:  data,
	})
}

// UPDATE PRODUCT
func UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedData models.Product

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	for i, product := range models.ListProduct {
		if product.ID == id {

			if updatedData.Picture != "" {
				models.ListProduct[i].Picture = updatedData.Picture
			}
			if updatedData.Name != "" {
				models.ListProduct[i].Name = updatedData.Name
			}
			if updatedData.Desc != "" {
				models.ListProduct[i].Desc = updatedData.Desc
			}
			if updatedData.Price != 0 {
				models.ListProduct[i].Price = updatedData.Price
			}
			models.ListProduct[i].FlashSale = updatedData.FlashSale
			if updatedData.Ratings != 0 {
				models.ListProduct[i].Ratings = updatedData.Ratings
			}
			if updatedData.Discount != 0 {
				models.ListProduct[i].Discount = updatedData.Discount
			}
			if updatedData.ReviewCounter != 0 {
				models.ListProduct[i].ReviewCounter = updatedData.ReviewCounter
			}
			if updatedData.Quantity != 0 {
				models.ListProduct[i].Quantity = updatedData.Quantity
			}
			if updatedData.Size != "" {
				models.ListProduct[i].Size = updatedData.Size
			}
			if updatedData.Variant != "" {
				models.ListProduct[i].Variant = updatedData.Variant
			}

			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Product updated successfully",
				Result:  models.ListProduct[i],
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "Product not found",
	})
}

// DELETE PRODUCT
func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, product := range models.ListProduct {
		if product.ID == id {
			models.ListProduct = append(models.ListProduct[:i], models.ListProduct[i+1:]...)

			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Product deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "Product not found",
	})
}