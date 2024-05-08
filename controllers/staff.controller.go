package controllers

import (
	"eniqilo_store/db"
	"eniqilo_store/helper"
	"eniqilo_store/helper/jwt"
	"eniqilo_store/models"
	"eniqilo_store/types"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type StaffController struct{}

func (h StaffController) Register(c *gin.Context) {
	db := db.GetDB()
	
	var request types.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errList := helper.ValidateRegisterRequest(&request)
	if len(errList) > 0 {
		errorMap := gin.H{
			"error": errList,
		}
		c.JSON(400, errorMap)
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM public.staff WHERE \"phoneNumber\" = $1", request.PhoneNumber).Scan(&count)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(409, gin.H{"error": "phone number exists"})
		return
	}

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), cost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Password = string(hashedPass)

	var staff models.Staff
	if err := db.QueryRow("INSERT INTO public.staff (\"phoneNumber\", name, password) VALUES ($1, $2, $3) RETURNING id, \"phoneNumber\", name", request.PhoneNumber, request.Name, request.Password).Scan(&staff.UserId, &staff.PhoneNumber, &staff.Name); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	
	token, err := jwt.SignJWT(staff.UserId, staff.PhoneNumber)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	data := map[string]string{
		"phoneNumber": staff.PhoneNumber,
		"name": staff.Name,
		"accessToken": token,
	}
	payload := gin.H{
		"message": "User registered successfully.",
		"data": data}
	c.JSON(201, payload)
}