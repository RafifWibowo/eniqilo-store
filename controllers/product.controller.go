package controllers

import (
	"eniqilo_store/db"
	"eniqilo_store/helper"
	"eniqilo_store/models"
	"eniqilo_store/types"
	"log"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (h ProductController) CreateProduct(c *gin.Context) {
	db := db.GetDB()

	var request types.ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	errList := helper.ValidateProductRequest(&request)
	if len(errList) > 0 {
		errorMap := gin.H{
			"error": errList,
		}
		c.JSON(400, errorMap)
		return
	}

	var product models.Product
	if err := db.QueryRow("INSERT INTO public.product (name, sku, category, \"imageUrl\", notes, price, stock, location, \"isAvailable\", \"createdAt\", \"updatedAt\") VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()) RETURNING id, \"createdAt\"", request.Name, request.SKU, request.Category, request.ImageUrl, request.Notes, request.Price, request.Stock, request.Location, request.IsAvailable).Scan(&product.UserId, &product.CreatedAt); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	createdAt, err := helper.ConvertToISO860(product.CreatedAt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	data := map[string]string{
		"id": product.UserId,
		"createdAt": createdAt,
	}
	payload := gin.H{
		"message": "success",
		"data": data,
	}
	c.JSON(201, payload)
}

func (h ProductController) UpdateProduct(c *gin.Context){
	db := db.GetDB()
	productId := c.Param("productId")

	var request types.ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	errList := helper.ValidateProductRequest(&request)
	if len(errList) > 0 {
		errorMap := gin.H{
			"error": errList,
		}
		c.JSON(400, errorMap)
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM public.product WHERE id = $1", productId).Scan(&count)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
	}

	_, err = db.Exec("UPDATE public.product SET name = $1, sku = $2, category = $3, \"imageUrl\" = $4, notes = $5, price = $6, stock = $7, location = $8, \"isAvailable\" = $9, \"updatedAt\" = NOW() WHERE id = $10", request.Name, request.SKU, request.Category, request.ImageUrl, request.Notes, request.Price, request.Stock, request.Location, request.IsAvailable, productId)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfully edit product.",
	})

}

func (h ProductController) SoftDeleteProduct(c *gin.Context) {
	db := db.GetDB()
	productId := c.Param("productId");

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM public.product WHERE id = $1", productId).Scan(&count)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		c.JSON(404, gin.H{"message": "Product not found"})
	}

	_, err = db.Exec("UPDATE public.product SET \"updatedAt\" = NOW(), \"deletedAt\" = NOW() WHERE id = $1", productId)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully delete product",
	})
}