package handlers

import (
	"net/http"

	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string              `json:"token"`
	User  models.UserResponse `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	if err := db.GetDB().Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !user.CheckPassword(req.Password) {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}
