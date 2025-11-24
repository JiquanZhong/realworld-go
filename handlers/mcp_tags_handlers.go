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
// @Tags mcp-tags
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param by query string false "排序字段" default(id)
// @Param asc query bool false "是否升序" default(true)
// @Success 200 {object} utils.Response "返回分页结果: total, page, list([]models.McpTagResponse)"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcp-tags [get]
func GetMcpTags(c *gin.Context) {
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	listOptions := utils.ListOptions{
		By:  c.DefaultQuery("by", "id"),
		Asc: c.DefaultQuery("asc", "true") == "true",
	}

	pagination, err := services.Services().McpTag.ListMcpTags(uint(page), uint(pageSize), listOptions)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, pagination)
}

// CreateMcpTag godoc
// @Summary 创建 MCP 标签
// @Description 创建 MCP 标签
// @Tags mcp-tags
// @Accept json
// @Produce json
// @Param request body services.McpTagRequest true "创建 MCP 标签请求"
// @Success 200 {object} utils.Response "返回创建结果: id, name"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcp-tags [post]
func CreateMcpTag(c *gin.Context) {
	var req services.McpTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mcpTagResponse, err := services.Services().McpTag.CreateMcpTag(req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, mcpTagResponse)
}

// UpdateMcpTag godoc
// @Summary 更新 MCP 标签
// @Description 更新 MCP 标签
// @Tags mcp-tags
// @Accept json
// @Produce json
// @Param id path int true "MCP 标签ID"
// @Param request body services.McpTagRequest true "更新 MCP 标签请求"
// @Success 200 {object} utils.Response "返回更新结果: id, name"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcp-tags/{id} [put]
func UpdateMcpTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req services.McpTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
	}

	mcpTagResponse, err := services.Services().McpTag.UpdateMcpTag(uint(id), req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, mcpTagResponse)
}

// DeleteMcpTag godoc
// @Summary 删除 MCP 标签
// @Description 删除 MCP 标签
// @Tags mcp-tags
// @Accept json
// @Produce json
// @Param id path int true "MCP 标签ID"
// @Success 200 {object} utils.Response "返回删除结果: id, name"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcp-tags/{id} [delete]
func DeleteMcpTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = services.Services().McpTag.DeleteMcpTag(uint(id))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, "MCP 标签删除成功")
}
