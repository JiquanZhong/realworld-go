package services

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
)

type McpTagService interface {
	ListMcpTags(page, pageSize uint, listOptions ListOptions) (utils.Pagination, error)
}

type mcpTagService struct{}

func NewMcpTagService() McpTagService {
	return &mcpTagService{}
}

func (s *mcpTagService) ListMcpTags(page, pageSize uint, listOptions ListOptions) (utils.Pagination, error) {

	var mcpTags []models.McpTag

	offset := (page - 1) * pageSize
	var total int64
	db.GetDB().Model(&models.McpTag{}).Count(&total)
	err := db.GetDB().Offset(int(offset)).Limit(int(pageSize)).Find(&mcpTags).Error
	if err != nil {
		return utils.Pagination{}, err
	}

	var mcpTagResponses []models.McpTagResponse
	for _, mcpTag := range mcpTags {
		mcpTagResponses = append(mcpTagResponses, mcpTag.ToResponse())
	}

	return utils.Pagination{
		Total:    uint(total),
		Page:     page,
		PageSize: pageSize,
		List:     &mcpTagResponses,
	}, nil

}
