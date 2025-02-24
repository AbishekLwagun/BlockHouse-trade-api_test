package handler

import (
	"BlockHouse_Test_AbishekLwgaun/db-connect"
	"BlockHouse_Test_AbishekLwgaun/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// for performing post request
func PostOrder(create *gin.Context) {
	var order models.Order

	if err := create.ShouldBindJSON(&order); err != nil {
		create.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingOrder models.Order
	if err := db_connect.DB.Where("symbol = ? AND price = ? AND order_type = ?", order.Symbol, order.Price, order.OrderType).First(&existingOrder).Error; err == nil {
		// If the order exists, return a conflict error
		create.JSON(http.StatusConflict, gin.H{"error": "Order already exists"})
		return
	}

	db_connect.DB.Create(&order)
	create.JSON(http.StatusCreated, order)
}

// for performing get request
func GetOrder(get *gin.Context) {
	var orders []models.Order
	db_connect.DB.Find(&orders)
	fmt.Println("Order saved")

	// test !
	//get.JSON(http.StatusOK, orders)

	prettyJSON, err := json.MarshalIndent(orders, "", "    ")
	if err != nil {
		get.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format JSON"})
		return
	}

	get.Data(http.StatusOK, "application/json", prettyJSON)
}
