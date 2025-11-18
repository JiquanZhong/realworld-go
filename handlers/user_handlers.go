package handlers

import (
	"net/http"
	"strconv"

	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.SetPassword(user.Password); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, user.ToResponse())
}

func GetUsers(c *gin.Context) {
	var users []models.User

	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var total int64
	db.GetDB().Model(&models.User{}).Count(&total)
	db.GetDB().Offset(offset).Limit(pageSize).Find(&users)

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	utils.Success(c, gin.H{
		"total": total,
		"page":  page,
		"list":  userResponses,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.GetDB().First(&user, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	utils.Success(c, user.ToResponse())
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.GetDB().First(&user, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.GetDB().Save(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, user.ToResponse())
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := db.GetDB().Delete(&models.User{}, id).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "User deleted"})
}
