package seats

import "gorm.io/gorm"

type Repository interface {
    GetByHallID(hallID uint) ([]Seat, error)
}

type seatRepository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &seatRepository{db: db}
}

func (r *seatRepository) GetByHallID(hallID uint) ([]Seat, error) {
    var seats []Seat
    if err := r.db.Where("hall_id = ?", hallID).Find(&seats).Error; err != nil {
        return nil, err
    }
    return seats, nil
}
