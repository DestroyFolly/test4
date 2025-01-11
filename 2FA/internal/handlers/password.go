package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"2FA/pkg/email"
)

func ResetPasswordHandler(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		OldPassword string `json:"oldPassword"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Email != adminEmail || request.OldPassword != adminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	code := generateVerificationCode()
	verificationCodes[request.Email] = code

	email.SendVerificationCode(code)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset code sent to email"})
}

func VerifyResetHandler(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		Code        string `json:"code"`
		NewPassword string `json:"newPassword"`
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

	adminPassword = request.NewPassword

	delete(verificationCodes, request.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully!"})
}
