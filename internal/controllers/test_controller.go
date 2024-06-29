package controllers

import (
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/models"
	"github.com/lewislewin/testify-webserver/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateTest(c *gin.Context) {
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateTest(&test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test created successfully"})
}

func GetTest(c *gin.Context) {
	id := c.Param("id")
	test, err := services.GetTest(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	c.JSON(http.StatusOK, test)
}

func UpdateTest(c *gin.Context) {
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateTest(id, &test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test updated successfully"})
}

func DeleteTest(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteTest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test deleted successfully"})
}

func ListTests(c *gin.Context) {
	tests, err := services.ListTests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list tests"})
		return
	}

	c.JSON(http.StatusOK, tests)
}
