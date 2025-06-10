package models

import (
	"github.com/google/uuid"
	"time"
)

// File represents a file stored in the database
type File struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar"`
	Data      []byte    `json:"data" gorm:"type:bytea"` // For PostgreSQL, use bytea
	Size      int64     `json:"size" gorm:"type:bigint"`
	MimeType  string    `json:"mime_type" gorm:"type:varchar"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
