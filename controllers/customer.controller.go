package controllers

import (
	"eniqilo_store/db"
	"eniqilo_store/helper"
	"eniqilo_store/models"
	"eniqilo_store/types"
	"log"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{}

func (h CustomerController) CustomerRegister(c *gin.Context) {
	db := db.GetDB()

	var request types.CustomerRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	errList := helper.ValidateCustomerRegisterRequest(&request)
	if len(errList) > 0 {
		errorMap := gin.H{
			"error": errList,
		}
		c.JSON(400, errorMap)
		return
	}

	checkInStaff, err := helper.CheckPhoneExist(db, "staff", request.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if checkInStaff {
		c.JSON(409, gin.H{"message": "Phone number already exists"})
		return
	}
	
	checkInCustomer, err := helper.CheckPhoneExist(db, "customer", request.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if checkInCustomer {
		c.JSON(409, gin.H{"message": "Phone number already exists"})
		return
	}

	var customer models.Customer
	if err := db.QueryRow("INSERT INTO public.customer (\"phoneNumber\", name, \"createdAt\", \"updatedAt\") VALUES ($1, $2, NOW(), NOW()) RETURNING id, \"phoneNumber\", name", request.PhoneNumber, request.Name).Scan(&customer.UserId, &customer.PhoneNumber, &customer.Name); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": err.Error()})
	}

	data := map[string]string{
		"userId": customer.UserId,
		"phoneNumber": customer.PhoneNumber,
		"name": customer.Name,
	}
	payload := gin.H{
		"message": "Customer registered successfully",
		"data": data,
	}
	c.JSON(201, payload)

}