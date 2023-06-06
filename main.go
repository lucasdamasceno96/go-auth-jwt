package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-auth-jwt.git/controllers"
	"github.com/lucasdamasceno96/go-auth-jwt.git/initializers"
	"github.com/lucasdamasceno96/go-auth-jwt.git/middleware"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	
  r := gin.Default()
  r.Use(cors.Default())
  r.POST("/signup", controllers.SignUp)
  r.POST("/login", controllers.Login)
  r.POST("/validate",middleware.RequireAuth, controllers.Validate)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
