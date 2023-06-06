package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-auth-jwt.git/controllers"
	"github.com/lucasdamasceno96/go-auth-jwt.git/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	
  r := gin.Default()
  r.POST("/signup", controllers.SignUp)
    r.POST("/login", controllers.Login)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
