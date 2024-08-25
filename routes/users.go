package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	fmt.Println(user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
