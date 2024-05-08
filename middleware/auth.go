package middleware

import (
	"database/sql"
	"eniqilo_store/db"
	"eniqilo_store/helper/jwt"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func getBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	
	return jwtToken[1], nil
}

func AuthMiddleware(c *gin.Context) {
	token, err := getBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": err.Error()})
		return
	}
	userId, err := jwt.ParseToken(token)
	if err != nil {
		log.Fatal(err.Error())
		c.AbortWithStatusJSON(401, gin.H{"message": err.Error()})
		return
	}

	db := db.GetDB()
	var id string
	err = db.QueryRow("SELECT id FROM public.staff WHERE id = $1 LIMIT 1", userId).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Staff not found"})
			return
		}
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Set("userId", id)
	c.Next()
}