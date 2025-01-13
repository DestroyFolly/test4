package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"2FA/pkg/email"
)

var (
	adminEmail        = os.Getenv("RECIPIENT_EMAIL_ADDRESS")
	adminPassword     = "123123"
	verificationCodes = make(map[string]string)
)


func LoginHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println(request.Email, adminEmail, request.Password, adminPassword)

	if request.Email != adminEmail || request.Password != adminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	code := generateVerificationCode()
	verificationCodes[request.Email] = code

	email.SendVerificationCode(code)

	c.JSON(http.StatusOK, gin.H{"message": "Verification code sent to email"})
}

func VerifyLoginHandler(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	expectedCode, exists := verificationCodes[request.Email]
	if !exists || expectedCode != request.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	delete(verificationCodes, request.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Auth success!"})
}

func generateVerificationCode() string {
	if os.Getenv("TEST") != "" {
		return "123456"
	}
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
