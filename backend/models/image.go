package models
import (
"time"
"github.com/google/uuid"
)
// Image representa a entidade de imagens
type Image struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar"`
	MimeType  string    `json:"mime_type" gorm:"type:varchar"`
	Data      []byte    `json:"data" gorm:"type:bytea"` // For PostgreSQL
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
