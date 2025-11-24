package seats

import "time"

type Seat struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
    HallID     uint      `gorm:"not null" json:"hall_id"`
    SeatNumber int       `gorm:"not null" json:"seat_number"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`

}

func (Seat) TableName() string {
	return "seats"
}