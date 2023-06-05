package initializers

import "github.com/lucasdamasceno96/go-auth-jwt.git/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
