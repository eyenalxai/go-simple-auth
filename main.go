package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-simple-auth/handlers"
	"go-simple-auth/middleware"
	"go-simple-auth/utils"
	"log"
)

func main() {
	r := gin.Default()

	jwtSecret, err := utils.GetJWTSecret()
	if err != nil {
		log.Println(err.Error())
		return
	}

	queries, err := utils.GetDbConnection()
	if err != nil {
		log.Println(err.Error())
		return
	}

	authorized := r.Group("/", middleware.AuthRequired(jwtSecret))
	{
		authorized.GET("/protected", handlers.ProtectedHandler)
	}
	r.POST("/register", handlers.RegisterHandler(queries, jwtSecret))
	r.POST("/login", handlers.LoginHandler(queries, jwtSecret))

	err = r.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
