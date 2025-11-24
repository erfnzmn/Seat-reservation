package shows

import "time"

type Show struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HallID    uint      `gorm:"not null" json:"hall_id"`
	MovieName string    `gorm:"type:varchar(255);not null" json:"movie_name"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Show) TableName() string {
	return "shows"
}