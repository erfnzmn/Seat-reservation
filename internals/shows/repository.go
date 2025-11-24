package shows

import (
	"gorm.io/gorm"
	"errors"
)

var (
	ErrShowNotFound = errors.New("show not found")
)

type Repository interface {
	GetAll() ([]Show, error)
	GetByID(id uint) (*Show, error)
	Create(show *Show) error
	Update(show *Show) error
	Delete(id uint) error
}

type showRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &showRepository{db: db}
}

func (r *showRepository) GetAll() ([]Show, error) {
	var shows []Show
	if err := r.db.Find(&shows).Error; err != nil {
		return nil, err
	}
	return shows, nil
}

func (r *showRepository) GetByID(id uint) (*Show, error) {
	var show Show
	if err := r.db.First(&show, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShowNotFound
		}
		return nil, err
	}
	return &show, nil
}

func (r *showRepository) Create(show *Show) error {
	return r.db.Create(show).Error
}

func (r *showRepository) Update(show *Show) error {
	// Ensure the row exists
	if _, err := r.GetByID(show.ID); err != nil {
		return err
	}
	return r.db.Save(show).Error
}

func (r *showRepository) Delete(id uint) error {
	result := r.db.Delete(&Show{}, id)
	if result.RowsAffected == 0 {
		return ErrShowNotFound
	}
	return result.Error
}