package reservation

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrReservationNotFound = errors.New("reservation not found")
)

type Repository interface {
	Create(res *Reservation) error
	GetByID(id uint) (*Reservation, error)
	GetBySeatAndShow(seatID uint, showID uint) (*Reservation, error)
	GetByShowID(showID uint) ([]Reservation, error)
	Cancel(id uint) error
}

type reservationRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) Create(res *Reservation) error {
	return r.db.Create(res).Error
}

func (r *reservationRepository) GetByID(id uint) (*Reservation, error) {
	var res Reservation
	if err := r.db.First(&res, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}
		return nil, err
	}
	return &res, nil
}

func (r *reservationRepository) GetBySeatAndShow(seatID uint, showID uint) (*Reservation, error) {
	var res Reservation
	if err := r.db.Where("seat_id = ? AND show_id = ? AND status = ?", seatID, showID, "confirmed").
		First(&res).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}
		return nil, err
	}
	return &res, nil
}

func (r *reservationRepository) GetByShowID(showID uint) ([]Reservation, error) {
	var list []Reservation
	if err := r.db.Where("show_id = ?", showID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *reservationRepository) Cancel(id uint) error {
	result := r.db.Model(&Reservation{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     "cancelled",
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrReservationNotFound
	}

	return nil
}
