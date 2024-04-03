package main

import (
	"rest-api/database"
	"rest-api/models"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.StartDB()

	router.POST("/orders", createOrderHandler)
    router.GET("/orders", getOrdersHandler)

    router.POST("/items", createItemHandler)
    router.GET("/items", getItemsHandler)

    router.Run(":8080")
}

func createOrderHandler(c *gin.Context) {
    // Parse form data
    customerName := c.PostForm("customer_name")
    if customerName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing customer_name parameter"})
        return
    }

    // Create the order
    err := createOrder(customerName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // Return success response
    c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func createOrder(customer_name string) error {
	db := database.GetDB()

	Order := models.Order {
		CustomerName: customer_name,
	}

	err := db.Create(&Order).Error

	return err
}

func getOrdersHandler(c *gin.Context) {
    db := database.GetDB()

    var orders []models.Order
    if err := db.Preload("Items").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

     // Convert orders to custom format
     var customOrders []gin.H
     for _, order := range orders {
         customOrder := gin.H{
             "OrderId":      order.OrderId,
             "CustomerName": order.CustomerName,
             "Items":        order.Items,
             "OrderedAt":    order.OrderedAt,
         }
         customOrders = append(customOrders, customOrder)
     }

    // Return the list of orders
    c.JSON(http.StatusOK, customOrders)
}

func createItemHandler(c *gin.Context) {
    // Parse the form data
    itemCode := c.PostForm("item_code")
    description := c.PostForm("description")
    quantity := c.PostForm("quantity")
    orderIDStr := c.PostForm("order_id") // Using OrderId here

    // Convert orderIDStr to uint
    orderId, err := strconv.ParseUint(orderIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OrderId"})
        return
    }

    // Create the item
    err = createItem(itemCode, description, quantity, uint(orderId))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
        return
    }

    // Return success response
    c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully"})
}

func createItem(itemCode, description, quantity string, orderId uint) error {
    db := database.GetDB()

    item := models.Item{
        ItemCode:    itemCode,
        Description: description,
        Quantity:    quantity,
        OrderId:     orderId,
    }

    if err := db.Create(&item).Error; err != nil {
        return err
    }

    return nil
}

func getItemsHandler(c *gin.Context) {
    db := database.GetDB()

    var items []models.Item
    if err := db.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
        return
    }

    // Return the list of items
    c.JSON(http.StatusOK, items)
}