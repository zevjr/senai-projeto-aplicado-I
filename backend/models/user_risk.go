package models
import (
"time"
"github.com/google/uuid"
)
// UserRisk representa a relação entre usuários e riscos
type UserRisk struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	RiskUID   uuid.UUID `json:"risk_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
