package controllers

import (
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/models"
	"github.com/lewislewin/testify-webserver/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateTestSuite(c *gin.Context) {
	var testSuite models.TestSuite
	if err := c.ShouldBindJSON(&testSuite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateTestSuite(&testSuite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create test suite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test suite created successfully"})
}

func GetTestSuite(c *gin.Context) {
	id := c.Param("id")
	testSuite, err := services.GetTestSuite(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test suite not found"})
		return
	}

	c.JSON(http.StatusOK, testSuite)
}

func UpdateTestSuite(c *gin.Context) {
	var testSuite models.TestSuite
	if err := c.ShouldBindJSON(&testSuite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateTestSuite(id, &testSuite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update test suite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test suite updated successfully"})
}

func DeleteTestSuite(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteTestSuite(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete test suite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test suite deleted successfully"})
}

func ListTestSuites(c *gin.Context) {
	testSuites, err := services.ListTestSuites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list test suites"})
		return
	}

	c.JSON(http.StatusOK, testSuites)
}
