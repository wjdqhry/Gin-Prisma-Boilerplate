package middleware

import (
	"Gin-Prisma-Boilerplate/db"
	"Gin-Prisma-Boilerplate/jwtAuth"
	"Gin-Prisma-Boilerplate/prisma"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORSmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}

func JWTmiddleware(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusBadRequest)

	}
	tokenString := strings.Split(authHeader, BEARER_SCHEMA)[1]

	if tokenString == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	token, err := jwtAuth.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if !token.Valid {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	claims := token.Claims.(jwt.MapClaims)
	for k, v := range claims {
		if k == "phoneNumber" {
			// user, err := prisma.Client.User.FindOne(
			// 	db.User.PhoneNumber.Equals(v.(string)),
			// ).Exec(prisma.Ctx)
			user, err := prisma.Client.User.FindOne(
				db.User.PhoneNumber.Equals(v.(string)),
			).Exec(prisma.Ctx)
			//log.Println(user[3])
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				log.Print(user)
				c.Set("user", user)
				c.Next()
			}

		}
	}
	//c.AbortWithStatus(http.StatusUnauthorized)
}
