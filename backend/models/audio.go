package models
import (
"time"
"github.com/google/uuid"
)
// Audio representa a entidade de áudios
type Audio struct {
	UID         uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar"`
	MimeType    string    `json:"mime_type" gorm:"type:varchar"`
	Data        []byte    `json:"data" gorm:"type:bytea"` // For PostgreSQL
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
	Transcricao string    `json:"transcricao,omitempty" gorm:"type:text"` // Optional transcription field
}
