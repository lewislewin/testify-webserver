package controllers

import (
	"net/http"
	"testify-webserver/models"
	"testify-webserver/services"

	"github.com/gin-gonic/gin"
)

func CreateEndpoint(c *gin.Context) {
	var endpoint models.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateEndpoint(&endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create endpoint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint created successfully"})
}

func GetEndpoint(c *gin.Context) {
	id := c.Param("id")
	endpoint, err := services.GetEndpoint(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func UpdateEndpoint(c *gin.Context) {
	var endpoint models.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateEndpoint(id, &endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update endpoint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint updated successfully"})
}

func DeleteEndpoint(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteEndpoint(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete endpoint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint deleted successfully"})
}

func ListEndpoints(c *gin.Context) {
	endpoints, err := services.ListEndpoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list endpoints"})
		return
	}

	c.JSON(http.StatusOK, endpoints)
}
