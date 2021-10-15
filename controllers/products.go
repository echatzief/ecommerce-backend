package controllers

import (
	"github.com/gin-gonic/gin"
)

//  GET /api/products
func findProducts(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"find products",
	})
}

//  GET /api/products/:id
func findOneProduct(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"find one product",
	})
}

//  PUT /api/products/:id
func updateOneProducts(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"find one product",
	})
}