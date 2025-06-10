package dto

import (
	"github.com/google/uuid"
	"time"
)

type ImageWithoutData struct {
	UID       uuid.UUID `json:"uid"`
	Name      string    `json:"name"`
	MimeType  string    `json:"mime_type"`
	CreatedAt time.Time `json:"created_at"`
}
