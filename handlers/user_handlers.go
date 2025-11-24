package handlers

import (
	"net/http"
	"strconv"

	"github.com/JiquanZhong/realworld-go/services"
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
	var req services.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := services.Services().User.CreateUser(req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
	}
	utils.Success(c, userResponse)
}

// GetUsers godoc
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param by query string false "排序字段" default(id)
// @Param asc query bool false "是否升序" default(true)
// @Success 200 {object} utils.Response "返回分页结果: total, page, list([]models.UserResponse)"
// @Success 404 {object} utils.Response "没有权限访问"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users [get]
func GetUsers(c *gin.Context) {

	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	listOptions := utils.ListOptions{
		By:  c.DefaultQuery("by", "id"),
		Asc: c.DefaultQuery("asc", "true") == "true",
	}

	pagination, err := services.Services().User.ListUsers(uint(page), uint(pageSize), listOptions)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, pagination)

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := services.Services().User.GetUser(uint(id))

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, userResponse)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req services.UpdateUserRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := services.Services().User.UpdateUser(id, req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, userResponse)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = services.Services().User.DeleteUser(uint(id))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "User deleted")
}

// UpdateUserRole godoc
// @Summary 修改用户角色
// @Description 管理员修改用户角色（仅管理员可访问）
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "用户ID"
// @Param request body services.UpdateUserRoleRequest true "角色信息"
// @Success 200 {object} utils.Response "返回更新后的用户数据 (models.UserResponse)"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 403 {object} utils.Response "权限不足"
// @Failure 404 {object} utils.Response "用户未找到"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /users/{id}/role [put]
func UpdateUserRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req services.UpdateUserRoleRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := services.Services().User.UpdateUserRole(uint(id), req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, userResponse)
}
