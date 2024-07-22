package controllers

import (
	"database/sql"
	"fmt"
	"rucara-services/database"
	"rucara-services/dtos"
	"rucara-services/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserSignUp handles user registration
func UserSignUp(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	_, err = database.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

// UserLogin handles user login and returns a JWT token
func UserLogin(c *fiber.Ctx) error {
	credentials := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	fmt.Println("Received login request for email: ", credentials.Email)
	fmt.Println("Received login request for email: ", credentials.Password)

	var user models.User
	row := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", credentials.Email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	fmt.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	fmt.Println(user)

	// Create JWT Token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &dtos.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "octagon",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(dtos.JWT_KEY)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"token": token})
}
