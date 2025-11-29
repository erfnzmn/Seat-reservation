package waitinglist

import "time"

type WaitingListStatus string

const (
	WaitingListStatusWaiting   WaitingListStatus = "waiting"
	WaitingListStatusAssigned  WaitingListStatus = "assigned"
	WaitingListStatusCancelled WaitingListStatus = "cancelled"
)

type WaitingList struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ShowID uint `gorm:"not null;index" json:"show_id"`
	// SeatID در ابتدا 0 است؛ وقتی worker صندلی را به این نفر اختصاص داد، مقدار می‌گیرد
	SeatID uint `gorm:"not null;default:0;index" json:"seat_id"`

	UserName  string `gorm:"size:100;not null" json:"user_name"`
	UserPhone string `gorm:"size:20;not null" json:"user_phone"`

	Status WaitingListStatus `gorm:"type:varchar(20);not null;default:'waiting'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (WaitingList) TableName() string {
	return "waiting_lists"
}
