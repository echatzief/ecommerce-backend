package controllers

import (
	"encoding/json"
	"fmt"
	"errors"
	"strings"
	"github.com/gin-gonic/gin"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"services/laiki-eu-backend/helpers"
	"services/laiki-eu-backend/models"
)

//  GET /api/users/count
func countUsers(ctx *gin.Context, conn *sql.DB) {

	// Parse the query string 
	queryString, err := parseUsersQuery(ctx, "SELECT COUNT(*) FROM Users")
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Do the query 
	rows, err := conn.Query(queryString)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Loop through the elements and count the items
	count := 0
	for rows.Next() { 
    err:= rows.Scan(&count)
		
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	
	ctx.JSON(200, gin.H{
		"count": count,
	})
	return
}

//  POST /api/users
func createUser(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"create user",
	})
}

//  DELETE /api/users
func deleteUsers(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"delete users",
	})
}

//  DELETE /api/users/:id
func deleteUser(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"delete a user",
	})
}

//  GET /api/users
func findUsers(ctx *gin.Context, conn *sql.DB){
	// Parse the query string 
	queryString, err := parseUsersQuery(ctx, "SELECT * FROM Users")
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Do the query 
	rows, err := conn.Query(queryString)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var users []models.User;

	for rows.Next() { 
		var user models.User;
    err:= rows.Scan(&user)
		
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		fmt.Println(user);
		users = append(users,user);
	}
	
	fmt.Println(users);
	usersJSON, err := json.Marshal(users)
	if err != nil {
			ctx.JSON(200, gin.H{
			"users": users,
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"users": usersJSON,
	})
	return
}

//  GET /api/users/:id
func findOneUser(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"find one user",
	})
}

//  PUT /api/users/:id
func updateOneUsers(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"update one user",
	})
}

//  GET /api/users/me
func findUserDetails(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message":"find user's details",
	})
}

func parseUsersQuery(ctx *gin.Context, query string) (string, error){

	if(len(ctx.Request.URL.Query()) > 0) {
		for k, q := range ctx.Request.URL.Query() 	{
			validStringFields := []string{"username","email","resetToken","confirmationToken","cellPhone","firstName","lastName","stripeAccountId"}
			validBoolFields := []string{"confirmed","blocked"}
	
			if helpers.StringArrayContains(validStringFields,k) || helpers.StringArrayContains(validBoolFields,k) {
				if strings.Contains(query,"WHERE")  {
					query = query + " AND " + k + "=\"" + q[0] + "\""
				} else {
					query = query + " WHERE " + k + "=\"" + q[0] + "\""
				}
			} else {
				 return "", errors.New("Query parameter '" + k + "' is not valid")
			}
		}
	}
	return query,nil
}