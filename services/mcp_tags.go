package services

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
)

type McpTagService interface {
	ListMcpTags(page, pageSize uint, listOptions utils.ListOptions) (utils.Pagination, error)
	CreateMcpTag(req McpTagRequest) (models.McpTagResponse, error)
	UpdateMcpTag(id uint, req McpTagRequest) (models.McpTagResponse, error)
	DeleteMcpTag(id uint) error
}

type mcpTagService struct{}

func NewMcpTagService() McpTagService {
	return &mcpTagService{}
}

type McpTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func (s *mcpTagService) ListMcpTags(page, pageSize uint, listOptions utils.ListOptions) (utils.Pagination, error) {

	var mcpTags []models.McpTag

	offset := (page - 1) * pageSize
	var total int64

	dbQuery := db.GetDB().Model(&models.McpTag{})
	if listOptions.By != "" {
		order := listOptions.By
		if !listOptions.Asc {
			order += " desc"
		}
		dbQuery = dbQuery.Order(order)
	}

	db.GetDB().Model(&models.McpTag{}).Count(&total)
	err := dbQuery.Offset(int(offset)).Limit(int(pageSize)).Find(&mcpTags).Error
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

func (s *mcpTagService) CreateMcpTag(req McpTagRequest) (models.McpTagResponse, error) {
	mcpTag := models.McpTag{
		Name: req.Name,
	}
	if err := db.GetDB().Create(&mcpTag).Error; err != nil {
		return models.McpTagResponse{}, err
	}
	return mcpTag.ToResponse(), nil
}

func (s *mcpTagService) UpdateMcpTag(id uint, req McpTagRequest) (models.McpTagResponse, error) {
	mcpTag := models.McpTag{
		ID:   id,
		Name: req.Name,
	}
	if err := db.GetDB().Save(&mcpTag).Error; err != nil {
		return models.McpTagResponse{}, err
	}
	return mcpTag.ToResponse(), nil
}

func (s *mcpTagService) DeleteMcpTag(id uint) error {
	if err := db.GetDB().Delete(&models.McpTag{}, id).Error; err != nil {
		return err
	}
	return nil
}
