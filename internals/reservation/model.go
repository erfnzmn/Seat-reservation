package reservation

import "time"

type Status string

const (
	StatusActive   Status = "active"
	StatusCanceled Status = "canceled"
)

type Reservation struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ShowID uint `gorm:"not null;index" json:"show_id"`
	HallID uint `gorm:"not null;index" json:"hall_id"`
	SeatID uint `gorm:"not null;index" json:"seat_id"`

	UserName  string `gorm:"size:100;not null" json:"user_name"`
	UserPhone string `gorm:"size:20;not null" json:"user_phone"`

	Status Status `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Reservation) TableName() string {
	return "reservations"
}
