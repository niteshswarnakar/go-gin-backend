package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex:idx_email"`
	Password  string    `json:"-"`
	IsAdmin   bool      `json:"is_admin"`
	Inactive  bool      `json:"inactive"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
