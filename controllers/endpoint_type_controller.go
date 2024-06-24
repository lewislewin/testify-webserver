package controllers

import (
	"net/http"
	"testify-webserver/models"
	"testify-webserver/services"

	"github.com/gin-gonic/gin"
)

func CreateEndpointType(c *gin.Context) {
	var endpointType models.EndpointType
	if err := c.ShouldBindJSON(&endpointType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateEndpointType(&endpointType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create endpoint type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint type created successfully"})
}

func GetEndpointType(c *gin.Context) {
	id := c.Param("id")
	endpointType, err := services.GetEndpointType(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint type not found"})
		return
	}

	c.JSON(http.StatusOK, endpointType)
}

func UpdateEndpointType(c *gin.Context) {
	var endpointType models.EndpointType
	if err := c.ShouldBindJSON(&endpointType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateEndpointType(id, &endpointType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update endpoint type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint type updated successfully"})
}

func DeleteEndpointType(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteEndpointType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete endpoint type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint type deleted successfully"})
}

func ListEndpointTypes(c *gin.Context) {
	endpointTypes, err := services.ListEndpointTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list endpoint types"})
		return
	}

	c.JSON(http.StatusOK, endpointTypes)
}
