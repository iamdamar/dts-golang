package middlewares

import (
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Authentication() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		verifyToken, err := helpers.VerifyToken(c)
// 		_ = verifyToken

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
// 				"error": "Unauthenticated",
// 				"message": err.Error(),
// 			})

// 			return
// 		}

// 		c.Set("userData", verifyToken)
// 		c.Next()
// 	}
// }

func Authentication() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the JWT token from the request header or cookie
        tokenString, err := extractTokenFromRequest(c)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error":   "Unauthorized",
                "message": "Missing or invalid token",
            })
            return
        }

        // Verify the JWT token
        claims, err := verifyToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error":   "Unauthorized",
                "message": "Invalid token",
            })
            return
        }

        // Set the user data in the context
        c.Set("userData", claims)

        // Continue handling the request
        c.Next()
    }
}
