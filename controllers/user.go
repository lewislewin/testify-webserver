package controllers

import (
	"database/sql"
	"testify-webserver/database"
	"testify-webserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO users (email, password, company_id) VALUES (?, ?, ?)"
	_, err := database.DB.Exec(query, user.Email, user.Password, user.CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func GetUser(c *gin.Context) {
	var user models.User
	query := "SELECT id, email, password, company_id FROM users WHERE id = ?"
	row := database.DB.QueryRow(query, c.Param("id"))
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CompanyID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE users SET email = ?, password = ?, company_id = ? WHERE id = ?"
	_, err := database.DB.Exec(query, user.Email, user.Password, user.CompanyID, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := database.DB.Exec(query, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func ListUsers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, email, password, company_id FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CompanyID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
