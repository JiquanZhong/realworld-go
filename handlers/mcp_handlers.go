package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JiquanZhong/realworld-go/services"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

// GetMcpServices godoc
// @Summary 获取 MCP 服务列表
// @Description 分页获取 MCP 服务列表，支持关键词搜索
// @Tags mcps
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param by query string false "排序字段" default(id)
// @Param asc query bool false "是否升序" default(true)
// @Param tags query string false "标签 ID 列表，逗号分隔"
// @Param search query string false "搜索关键词（匹配名称、描述、分类）"
// @Success 200 {object} utils.Response "返回分页结果: total, page, list([]models.McpResponse)"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps [get]
func GetMcpServices(c *gin.Context) {
	page, _ := strconv.Atoi((c.DefaultQuery("page", "1")))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	searchKeyword := c.Query("search")
	listOptions := utils.ListOptions{
		By:  c.DefaultQuery("by", "id"),
		Asc: c.DefaultQuery("asc", "true") == "true",
	}

	var tagIds []uint
	if tagsParam := c.Query("tags"); tagsParam != "" {
		tagsStrs := strings.Split(tagsParam, ",")
		for _, tagStr := range tagsStrs {
			tagId, err := strconv.Atoi(tagStr)
			if err == nil {
				tagIds = append(tagIds, uint(tagId))
			} else {
				utils.Error(c, http.StatusBadRequest, "invalid tag id: "+tagStr)
				return
			}
		}
	}

	pagination, err := services.Services().Mcp.ListMcpServices(uint(page), uint(pageSize), listOptions, tagIds, searchKeyword)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, pagination)
}

// RegisterMcpService godoc
// @Summary 创建 MCP 服务
// @Description 注册一个新的 MCP 服务，并绑定标签
// @Tags mcps
// @Accept json
// @Produce json
// @Param request body services.RegisterMcpServiceRequest true "MCP 服务信息"
// @Success 200 {object} utils.Response "返回创建成功的 MCP 服务"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps [post]
func RegisterMcpService(c *gin.Context) {
	var req services.RegisterMcpServiceRequest
	userId, ok := utils.GetUserID(c)
	if !ok {
		utils.Error(c, http.StatusUnauthorized, "请先登录")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	req.SubmitterID = userId

	mcp, err := services.Services().Mcp.RegisterMcpService(req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, mcp.ToResponse())
}

// GetMcpService godoc
// @Summary 获取 MCP 服务详情
// @Description 根据 ID 获取 MCP 服务信息（含标签）
// @Tags mcps
// @Produce json
// @Param id path int true "MCP 服务 ID"
// @Success 200 {object} utils.Response "返回 MCP 服务详情"
// @Failure 404 {object} utils.Response "未找到 MCP 服务"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps/{id} [get]
func GetMcpService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	userId, _ := utils.GetUserID(c)

	mcpDetail, err := services.Services().Mcp.GetMcpService(uint(id), userId)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, mcpDetail)
}

// DeleteMcpService godoc
// @Summary 删除 MCP 服务
// @Description 根据 ID 删除 MCP 服务
// @Tags mcps
// @Produce json
// @Param id path int true "MCP 服务 ID"
// @Success 200 {object} utils.Response "删除成功"
// @Failure 404 {object} utils.Response "未找到 MCP 服务"
// @Failure 500 {object} utils.Response "服务器内部错误"
// @Router /mcps/{id} [delete]
func DeleteMcpService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := services.Services().Mcp.DeleteMcpService(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "deleted"})
}
