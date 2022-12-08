package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		result gin.H
	)
	result = gin.H{
		"result" : "haiiii sukses",
		"count" : 1,
	}
	c.JSON(http.StatusOK, result)
}
