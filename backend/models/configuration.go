package models
import (
"time"
"github.com/google/uuid"
)
// Configuration representa configurações de usuário
type Configuration struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Config    []byte    `json:"config" gorm:"type:jsonb"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
