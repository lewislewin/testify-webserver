package controllers

import (
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/models"
	"github.com/lewislewin/testify-webserver/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateCompany(&company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company created successfully"})
}

func GetCompany(c *gin.Context) {
	id := c.Param("id")
	company, err := services.GetCompany(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func UpdateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := services.UpdateCompany(id, &company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteCompany(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

func ListCompanies(c *gin.Context) {
	companies, err := services.ListCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}
