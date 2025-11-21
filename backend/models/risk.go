package models
import (
"time"
"github.com/google/uuid"
)
// Risk representa a entidade de riscos
type Risk struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Details   string    `json:"details" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
