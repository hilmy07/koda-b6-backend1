package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  any `json:"results"`
}

type Users struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ListUser []Users
var userCounter int

func main() {
	r := gin.Default()

	// GET ALL USERS
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "List of users",
			Result:  ListUser,
		})
	})

	// GET USER BY ID
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, user := range ListUser {
			if user.ID == id {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User found",
					Result:  user,
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// CREATE (REGISTER)
	r.POST("/register", func(ctx *gin.Context) {
		var data Users

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		userCounter++
		data.ID = strconv.Itoa(userCounter)

		ListUser = append(ListUser, data)

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register successfully",
			Result:  data,
		})
	})

	// UPDATE USER
	r.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var updatedData Users

		if err := ctx.ShouldBindJSON(&updatedData); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		for i, user := range ListUser {
			if user.ID == id {
				ListUser[i].Email = updatedData.Email
				ListUser[i].Password = updatedData.Password

				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User updated successfully",
					Result:  ListUser[i],
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// DELETE USER
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for i, user := range ListUser {
			if user.ID == id {
				ListUser = append(ListUser[:i], ListUser[i+1:]...)

				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User deleted successfully",
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
	})

	// LOGIN
	r.POST("/login", func(ctx *gin.Context) {
		var loginData Users

		if err := ctx.ShouldBindJSON(&loginData); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		for _, user := range ListUser {
			if user.Email == loginData.Email && user.Password == loginData.Password {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Login successfully",
					Result:  user,
				})
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Login failed",
		})
	})

	r.Run("localhost:8000")
}