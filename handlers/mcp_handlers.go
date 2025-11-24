package handlers

import (
	"net/http"
	"strconv"

	"github.com/JiquanZhong/realworld-go/services"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

// GetMcpServices godoc
// @Summary 获取 MCP 服务列表
// @Description 分页获取 MCP 服务列表
// @Tags mcps
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.Response "返回分页结果: total, page, list([]models.McpResponse)"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps [get]
func GetMcpServices(c *gin.Context) {
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pagination, err := services.Services().Mcp.ListMcpServices(uint(page), uint(pageSize), services.ListOptions{})

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, pagination)
}