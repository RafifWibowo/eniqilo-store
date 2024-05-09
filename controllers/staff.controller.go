package controllers

import (
	"database/sql"
	"eniqilo_store/db"
	"eniqilo_store/helper"
	"eniqilo_store/helper/hash"
	"eniqilo_store/helper/jwt"
	"eniqilo_store/models"
	"eniqilo_store/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	hashedPass, err := hash.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Password = hashedPass

	var staff models.Staff
	if err := db.QueryRow("INSERT INTO public.staff (\"phoneNumber\", name, password, \"createdAt\", \"updatedAt\") VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, \"phoneNumber\", name", request.PhoneNumber, request.Name, request.Password).Scan(&staff.UserId, &staff.PhoneNumber, &staff.Name); err != nil {
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
		"message": "Staff registered successfully.",
		"data": data}
	c.JSON(201, payload)
}

func (h StaffController) Login (c *gin.Context) {
	db := db.GetDB()

	var request types.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errList := helper.ValidateLoginRequest(&request)
	if len(errList) > 0 {
		errorMap := gin.H{
			"error": errList,
		}
		c.JSON(400, errorMap)
		return
	}

	var staff models.Staff
	err := db.QueryRow("SELECT id, \"phoneNumber\", name, password FROM staff WHERE \"phoneNumber\" = $1", request.PhoneNumber).Scan(&staff.UserId, &staff.PhoneNumber, &staff.Name, &staff.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Staff not found"})
			return
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !hash.CheckPassword(request.Password, staff.Password) {
		c.JSON(400, gin.H{"error": "Invalid password."})
	}

	token, err := jwt.SignJWT(staff.UserId, staff.PhoneNumber)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	data := map[string]string{
		"userId": staff.UserId,
		"phoneNumber": staff.PhoneNumber,
		"name": staff.Name,
		"accessToken": token,
	}

	payload := gin.H{
		"message": "Staff logged in successfully.",
		"data": data,
	}
	c.JSON(200, payload)
}