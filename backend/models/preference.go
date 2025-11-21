package models
import (
"time"
"github.com/google/uuid"
)
// Preference representa as preferências de usuário
type Preference struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Prefer    []byte    `json:"prefer" gorm:"type:jsonb"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
