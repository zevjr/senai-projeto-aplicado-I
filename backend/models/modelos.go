package models

import (
	"time"

	"github.com/google/uuid"
)

// User representa a entidade de usuários
type User struct {
	UID       uuid.UUID  `json:"uid" gorm:"type:uuid;primaryKey"`
	Username  string     `json:"username" gorm:"type:varchar"`
	Role      string     `json:"role" gorm:"type:varchar"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:timestamp"`
	Password  string     `json:"password,omitempty" gorm:"type:varchar"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar"`
}

// Risk representa a entidade de riscos
type Risk struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Details   string    `json:"details" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// UserRisk representa a relação entre usuários e riscos
type UserRisk struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	RiskUID   uuid.UUID `json:"risk_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// Register representa a entidade de registros
type Register struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar"`
	Body      string    `json:"body" gorm:"type:text"`
	RiskScale int       `json:"riskScale" gorm:"type:integer;not null"`
	Local     string    `json:"local" gorm:"type:varchar"`
	Status    string    `json:"status" gorm:"type:varchar"`
	ImageUID  uuid.UUID `json:"imageUid" gorm:"type:uuid"`
	AudioUID  uuid.UUID `json:"audioUid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// UserRegister representa a relação entre usuários e registros
type UserRegister struct {
	UID         uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	UserUID     uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	RegisterUID uuid.UUID `json:"register_uid" gorm:"type:uuid"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
}

// Image representa a entidade de imagens
type Image struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar"`
	MimeType  string    `json:"mime_type" gorm:"type:varchar"`
	Data      []byte    `json:"data" gorm:"type:bytea"` // For PostgreSQL
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// Audio representa a entidade de áudios
type Audio struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar"`
	MimeType  string    `json:"mime_type" gorm:"type:varchar"`
	Data      []byte    `json:"data" gorm:"type:bytea"` // For PostgreSQL
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// Preference representa as preferências de usuário
type Preference struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Prefer    []byte    `json:"prefer" gorm:"type:jsonb"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

// Configuration representa configurações de usuário
type Configuration struct {
	UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
	Config    []byte    `json:"config" gorm:"type:jsonb"`
	UserUID   uuid.UUID `json:"user_uid" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
