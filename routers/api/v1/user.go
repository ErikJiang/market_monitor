package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"user": "user", "value": "value"})

}

func AddUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

func EditUser(c *gin.Context) {

}

func RemoveUser(c *gin.Context) {

}