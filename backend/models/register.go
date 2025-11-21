package models
import (
"time"
"github.com/google/uuid"
)
// Register representa a entidade de registros
type Register struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar"`
	Body      string    `json:"body" gorm:"type:text"`
	RiskScale int       `json:"riskScale" gorm:"type:integer;"`
	Local     string    `json:"local" gorm:"type:varchar"`
	Status    string    `json:"status" gorm:"type:varchar"`
	ImageUID  uuid.UUID `json:"image_uid" gorm:"type:uuid"`
	AudioUID  uuid.UUID `json:"audio_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
