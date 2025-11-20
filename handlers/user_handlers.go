package handlers

import (
	"net/http"
	"strconv"

	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary 创建用户
// @Description 使用提供的信息创建新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "用户负载"
// @Success 200 {object} utils.Response "成功: 返回创建的用户数据 (models.UserResponse)"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users [post]
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

// GetUsers godoc
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.Response "返回分页结果: total, page, list([]models.UserResponse)"
// @Success 404 {object} utils.Response "没有权限访问"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users [get]
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

// GetUser godoc
// @Summary 通过ID获取用户
// @Description 根据用户ID获取用户详情
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "用户ID"
// @Success 200 {object} utils.Response "返回用户数据 (models.UserResponse)"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 404 {object} utils.Response "用户未找到"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.GetDB().First(&user, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	utils.Success(c, user.ToResponse())
}

// UpdateUser godoc
// @Summary 更新用户
// @Description 根据ID更新用户信息
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "用户ID"
// @Param user body models.User true "用户负载"
// @Success 200 {object} utils.Response "返回更新后的用户数据 (models.UserResponse)"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 404 {object} utils.Response "用户未找到"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users/{id} [put]
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

// DeleteUser godoc
// @Summary 删除用户
// @Description 根据ID删除用户
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "用户ID"
// @Success 200 {object} utils.Response "删除成功"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := db.GetDB().Delete(&models.User{}, id).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "User deleted"})
}
