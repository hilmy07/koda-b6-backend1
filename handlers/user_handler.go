package handlers

import (
	"backend1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  any 		`json:"results"`
}

// Get All Users godoc
// @Summary Get all users
// @Tags Users
// @Produce json
// @Success 200 {object} Response
// @Router /users [get]
func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "List of users",
		Result:  models.ListUser,
	})
}

// Get User By ID godoc
// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id} [get]
func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, user := range models.ListUser {
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
}

// Register godoc
// @Summary Register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.Users true "Register Data"
// @Success 200 {object} Response
// @Router /register [post]
func Register(ctx *gin.Context) {
	var data models.Users

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	cfg := argon2.DefaultConfig()
	hash, err := cfg.HashEncoded([]byte(data.Password))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	data.Password = string(hash)

	models.UserCounter++
	data.ID = strconv.Itoa(models.UserCounter)

	models.ListUser = append(models.ListUser, data)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Register successfully",
		Result:  data,
	})
}

// Update User godoc
// @Summary Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body models.Users true "Update Data"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id} [put]
func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedData models.Users

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	for i, user := range models.ListUser {
		if user.ID == id {

			// Update email kalau diisi
			if updatedData.Email != "" {
				models.ListUser[i].Email = updatedData.Email
			}

			// 🔐 Kalau password diisi → hash dulu
			if updatedData.Password != "" {
				cfg := argon2.DefaultConfig()
				hash, err := cfg.HashEncoded([]byte(updatedData.Password))
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, Response{
						Success: false,
						Message: "Failed to hash password",
					})
					return
				}
				models.ListUser[i].Password = string(hash)
			}

			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "User updated successfully",
				Result:  models.ListUser[i],
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: "User not found",
	})
}

// Delete User godoc
// @Summary Delete user
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, user := range models.ListUser {
		if user.ID == id {
			models.ListUser = append(models.ListUser[:i], models.ListUser[i+1:]...)

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
}

// Login godoc
// @Summary Login with registered account
// @Description Login using email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.Users true "Login Data"
// @Success 200 {object} Response
// @Router /login [post]
func Login(ctx *gin.Context) {
	var loginData models.Users

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	for _, user := range models.ListUser {
		if user.Email == loginData.Email {

			// 🔎 VERIFY PASSWORD
			ok, err := argon2.VerifyEncoded(
				[]byte(loginData.Password),
				[]byte(user.Password),
			)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, Response{
					Success: false,
					Message: "Password verification failed",
				})
				return
			}

			if ok {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Login successfully",
					Result:  user,
				})
				return
			}
		}
	}


	ctx.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: "Login failed",
	})
}