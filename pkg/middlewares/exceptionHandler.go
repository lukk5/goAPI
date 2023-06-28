package middlewares

import (
	"goAPI/pkg/exceptions"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Exception occurred:", err)

			switch e := err.(type) {
			case exceptions.BadRequestError:
				c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			case exceptions.NotFoundError:
				c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}

			c.Abort()
		}
	}()
	// Call the next handler
	c.Next()
}
