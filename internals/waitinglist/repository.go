package waitinglist

import "gorm.io/gorm"

type Repository interface {
	Add(entry *WaitingList) error
	GetNextInQueue(showID uint) (*WaitingList, error)
	MarkAsAssigned(id uint, seatID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Add(entry *WaitingList) error {
	return r.db.Create(entry).Error
}

func (r *repository) GetNextInQueue(showID uint) (*WaitingList, error) {
	var w WaitingList
	err := r.db.
		Where("show_id = ? AND status = ?", showID, WaitingListStatusWaiting).
		Order("created_at ASC").
		First(&w).Error

	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *repository) MarkAsAssigned(id uint, seatID uint) error {
	return r.db.
		Model(&WaitingList{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":  WaitingListStatusAssigned,
			"seat_id": seatID,
		}).Error
}
