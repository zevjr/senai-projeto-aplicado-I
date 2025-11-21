package models
import (
"time"
"github.com/google/uuid"
)
// UserRegister representa a relação entre usuários e registros
type UserRegister struct {
	UID         uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	UserUID     uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	RegisterUID uuid.UUID `json:"register_uid" gorm:"type:uuid"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
}
