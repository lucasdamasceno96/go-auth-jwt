package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-auth-jwt.git/initializers"
	"github.com/lucasdamasceno96/go-auth-jwt.git/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to read the body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to read the body",
		})
		return 
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to create user",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"User" : "Succesfully created",
	})

}
