package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/auth"
	"github.com/lewislewin/testify-webserver/internal/database"
	"github.com/lewislewin/testify-webserver/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Registration controller
// This function is to register a new USER and COMPANY, users can be added to a company once created by a company admin
func Register(c *gin.Context) {
	log.Println("Register called")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = string(hashedPassword)

	input.ID = uuid.New()

	query := "INSERT INTO users (id, email, first_name, surname, password, company_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = database.DB.Exec(query, input.ID, input.Email, input.FirstName, input.Surname, input.Password, input.CompanyID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	log.Println("User registered successfully")
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// This function is for any user to login, not just a company
func Login(c *gin.Context) {
	log.Println("Login called")
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = ?"
	row := database.DB.QueryRow(query, input.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		log.Printf("Invalid login attempt for email: %s", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateJWT(user.InternalID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	log.Println("Login successful")
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func GetCurrentUser(c *gin.Context) {
	log.Println("GetCurrentUser called")
	userID, exists := c.Get("userID")
	if !exists {
		log.Println("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	query := "SELECT id, email, company_id FROM users WHERE id = ?"
	row := database.DB.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.CompanyID)
	if err == sql.ErrNoRows {
		log.Println("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		log.Printf("Error retrieving user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
		return
	}

	log.Println("User retrieved successfully")
	c.JSON(http.StatusOK, user)
}
