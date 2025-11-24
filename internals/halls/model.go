package halls

import "time"


type Hall struct{
	ID         uint      `gorm:"primaryKey" json:"id"`
    Name       string    `gorm:"size:100;not null" json:"name"`
    TotalSeats int       `gorm:"not null" json:"total_seats"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

func (Hall) TableName() string {
	return "halls"
}
