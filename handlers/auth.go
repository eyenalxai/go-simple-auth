package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	users "go-simple-auth/postgresql"
	"go-simple-auth/types"
	"go-simple-auth/utils"
	"log"
	"net/http"
	"time"
)

func RegisterHandler(q *users.Queries, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cred types.Credentials

		if err := c.BindJSON(&cred); err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{"error": "credentials oopsie"})
			return
		}

		hashedPassword, err := utils.HashPassword(cred.Password)

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "password hashing oopsie"})
			return
		}

		uuidUser := uuid.New()
		params := users.CreateUserParams{
			ID:           uuidUser,
			Username:     cred.Username,
			PasswordHash: hashedPassword,
		}

		_, err = q.CreateUser(c, params)

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "user creation oopsie"})
			return
		}

		tokenString, err := utils.GetToken(cred.Username, jwtSecret)

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "token creation oopsie"})
			return
		}

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "token creation oopsie"})
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "session",
			Value:    tokenString,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		c.JSON(201, gin.H{"status": "user created"})
	}
}
func LoginHandler(q *users.Queries, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cred types.Credentials

		if err := c.BindJSON(&cred); err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{"error": "invalid credentials"})
			return
		}

		user, err := q.GetUser(c, cred.Username)

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "failed to retrieve user"})
			return
		}

		if !utils.CheckPasswordHash(cred.Password, user.PasswordHash) {
			log.Println("invalid password")
			c.JSON(403, gin.H{"error": "wrong password"})
			return
		}

		tokenString, err := utils.GetToken(cred.Username, jwtSecret)

		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{"error": "failed to generate token"})
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "session",
			Value:    tokenString,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		c.JSON(200, gin.H{"status": "login successful"})
	}
}
