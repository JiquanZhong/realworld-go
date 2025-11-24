package services

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
)

type McpService interface {
	// RegisterMcpService(req RegisterMcpServiceRequest) (models.McpService, error)
	ListMcpServices(page, pageSize uint, listOptions ListOptions) (utils.Pagination, error)
	// GetMcpService(id uint) (models.McpService, error)
}

type mcpService struct{}

func NewMcpService() McpService {
	return &mcpService{}
}

type ListOptions struct {
}

func (s *mcpService) ListMcpServices(page, pageSize uint, listOptions ListOptions) (utils.Pagination, error) {

	var mcpServices []models.McpService

	offset := (page - 1) * pageSize
	var total int64
	db.GetDB().Model(&models.McpService{}).Count(&total)
	err := db.GetDB().Preload("Tags").Offset(int(offset)).Limit(int(pageSize)).Find(&mcpServices).Error
	if err != nil {
		return utils.Pagination{}, err
	}

	var mcpResponses []models.McpResponse
	for _, mcpService := range mcpServices {
		mcpResponses = append(mcpResponses, mcpService.ToResponse())
	}

	return utils.Pagination{
		Total:    uint(total),
		Page:     page,
		PageSize: pageSize,
		List:     &mcpResponses,
	}, nil

}
