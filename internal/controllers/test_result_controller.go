package controllers

import (
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/models"
	"github.com/lewislewin/testify-webserver/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateTestResult(c *gin.Context) {
	var testResult models.TestResult
	if err := c.ShouldBindJSON(&testResult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateTestResult(&testResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create test result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test result created successfully"})
}

func GetTestResult(c *gin.Context) {
	id := c.Param("id")
	testResult, err := services.GetTestResult(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test result not found"})
		return
	}

	c.JSON(http.StatusOK, testResult)
}

func UpdateTestResult(c *gin.Context) {
	var testResult models.TestResult
	if err := c.ShouldBindJSON(&testResult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateTestResult(id, &testResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update test result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test result updated successfully"})
}

func DeleteTestResult(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteTestResult(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete test result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test result deleted successfully"})
}

func ListTestResults(c *gin.Context) {
	testResults, err := services.ListTestResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list test results"})
		return
	}

	c.JSON(http.StatusOK, testResults)
}
