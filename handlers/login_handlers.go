package handlers

import (
	"net/http"

	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/services"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type LoginResponse struct {
	Token string              `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  models.UserResponse `json:"user"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Successfully logged in"
// @Failure 400 {object} utils.Response "Invalid request body"
// @Failure 401 {object} utils.Response "Invalid email or password"
// @Failure 500 {object} utils.Response "Failed to generate token"
// @Router /users/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GetLogger().Error("Invalid request body", zap.Error(err))
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	loginResponse, err := services.Services().User.Login(req.Email, req.Password)
	if err != nil {
		utils.GetLogger().Error("Failed to login", zap.Error(err))
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	utils.Success(c, loginResponse)
}
