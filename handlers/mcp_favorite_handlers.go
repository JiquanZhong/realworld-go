package handlers

import (
	"net/http"
	"strconv"

	"github.com/JiquanZhong/realworld-go/services"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

// AddMcpServiceFavorite godoc
// @Summary 为 MCP 服务添加收藏
// @Description 用户为指定的 MCP 服务添加收藏
// @Tags mcps
// @Accept json
// @Produce json
// @Param id path int true "Mcp 服务 ID"
// @Success 200 {object} utils.Response "收藏成功"
// @Failure 400 {object} utils.Response "无效的 MCP 服务 ID"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps/{id}/favorite [post]
func AddMcpServiceFavorite(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的Mcp ID")
		return
	}

	userId, exists := utils.GetUserID(c)
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "请先登录")
		return
	}

	err = services.Services().McpFavorite.AddFavorite(uint(id), userId)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "已添加到收藏")
}

// RemoveMcpServiceFavorite godoc
// @Summary 取消 MCP 服务收藏
// @Description 用户取消对指定 MCP 服务的收藏
// @Tags mcps
// @Accept json
// @Produce json
// @Param id path int true "Mcp 服务 ID"
// @Success 200 {object} utils.Response "取消收藏成功"
// @Failure 400 {object} utils.Response "无效的 MCP 服务 ID"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps/{id}/favorite [delete]
func RemoveMcpServiceFavorite(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的Mcp ID")
		return
	}

	userId, exists := utils.GetUserID(c)
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "请先登录")
		return
	}

	err = services.Services().McpFavorite.RemoveFavorite(uint(id), userId)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, "已取消收藏")
}
