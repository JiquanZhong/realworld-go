package services

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/models"
)

type McpFavoriteService interface {
	AddFavorite(id uint, userId uint) error
	RemoveFavorite(id uint, userId uint) error
	GetFavoriteNum(id uint) (uint, error)
}

type mcpFavoriteService struct{}

func NewMcpFavoriteService() McpFavoriteService {
	return &mcpFavoriteService{}
}

func (s *mcpFavoriteService) AddFavorite(id uint, userId uint) error {
	var count int64
	err := db.GetDB().Model(&models.McpFavorite{}).
		Where("user_id = ? AND mcp_service_id = ?", userId, id).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count == 1 {
		return nil
	}

	favorite := models.McpFavorite{
		McpServiceID: id,
		UserID:       userId,
	}
	return db.GetDB().Create(&favorite).Error
}

func (s *mcpFavoriteService) RemoveFavorite(id uint, userId uint) error {
	var count int64
	err := db.GetDB().Model(&models.McpFavorite{}).
		Where("user_id = ? AND mcp_service_id = ?", userId, id).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	}

	return db.GetDB().Where("mcp_service_id = ? AND user_id = ?", id, userId).Delete(&models.McpFavorite{}).Error
}

func (s *mcpFavoriteService) GetFavoriteNum(id uint) (uint, error) {
	var count int64
	err := db.GetDB().Model(&models.McpFavorite{}).Where("mcp_service_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return uint(count), nil
}
