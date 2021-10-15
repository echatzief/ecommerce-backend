package controllers

import (
	"github.com/gin-gonic/gin"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func Setup(router *gin.Engine, conn *sql.DB){

	// Users 
	router.GET("/api/users/count", func (ctx *gin.Context){countUsers(ctx, conn)})
	router.POST("/api/users", createUser)
	router.DELETE("/api/users", deleteUsers)
	router.DELETE("/api/users/:id", deleteUser)
	router.GET("/api/users", func (ctx *gin.Context){findUsers(ctx, conn)})
	router.GET("/api/users/:id", findOneUser)
	router.GET("/api/users/me", findUserDetails)
	router.PUT("/api/users/:id", updateOneUsers)


  router.GET("/api/products", findProducts)
	router.GET("/api/products/:id", findOneProduct)
	router.PUT("/api/products/:id", updateOneProducts)
}
