package models

type ForumBoards struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Slug        string `gorm:"type:varchar(100);not null;unique" json:"slug"`
	Description string `gorm:"type:text" json:"description"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`
	IsHidden    bool   `gorm:"default:false" json:"is_hidden"`
}

func (ForumBoards) TableName() string {
	return "forum_boards"
}
