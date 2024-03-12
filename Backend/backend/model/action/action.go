package action

import db "backend/database"

type UserActionModel struct {
	UserID     int    `json:"user_id,omitempty" gorm:"column:user_id;" binding:"required"`
	ResourceID int    `json:"resource_id,omitempty" gorm:"column:resource_id;" binding:"required"`
	ActionType string `json:"action_type,omitempty" gorm:"column:action_type;" binding:"required"`
}

func (ua *UserActionModel) TableName() string {
	return "user_actions"
}

func (ua *UserActionModel) Create() error {
	return db.DB.
		Table(ua.TableName()).
		Create(ua).Error
}
