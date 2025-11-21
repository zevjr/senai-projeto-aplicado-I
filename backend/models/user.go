package models
import (
"time"
"github.com/google/uuid"
)
// User representa a entidade de usuários
type User struct {
	UID       uuid.UUID  `json:"uid" gorm:"type:uuid;primaryKey"`
	Username  string     `json:"username" gorm:"type:varchar"`
	Role      string     `json:"role" gorm:"type:varchar"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:timestamp"`
	Password  string     `json:"password,omitempty" gorm:"type:varchar"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar"`
}
