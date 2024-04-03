package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H {
		"id":			User.ID,
		"email":		User.Email,
		"username":		User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Unauthorized",
			"message": "invalid email/password",
		})

		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Unauthorized",
			"message": "invalid email/password",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H {
		"token": token,
	})
	
}

func UserUpdate(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)
    _, _ = db, contentType
    User := models.User{}

    // Extract user ID from the request URL parameters
    userIDStr := c.Param("userID")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": "Invalid user ID parameter",
        })
        return
    }
    User.ID = uint(userID)

    // Bind user data from request body
    if contentType == appJSON {
        c.ShouldBindJSON(&User)
    } else {
        c.ShouldBind(&User)
    }

    // Update user data in the database
    err = db.Debug().Save(&User).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    // Return updated user data
    c.JSON(http.StatusOK, gin.H{
        "id":       User.ID,
        "email":    User.Email,
        "username": User.Username,
    })
}
