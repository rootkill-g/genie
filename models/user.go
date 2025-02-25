package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"uniqueIndex"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
