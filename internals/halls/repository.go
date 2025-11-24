package halls

import "gorm.io/gorm"

type Repository interface {
    GetAll() ([]Hall, error)
    GetByID(id uint) (*Hall, error)
}

type hallRepository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &hallRepository{db: db}
}

func (r *hallRepository) GetAll() ([]Hall, error) {
    var halls []Hall
    if err := r.db.Find(&halls).Error; err != nil {
        return nil, err
    }
    return halls, nil
}

func (r *hallRepository) GetByID(id uint) (*Hall, error) {
    var hall Hall
    if err := r.db.First(&hall, id).Error; err != nil {
        return nil, err
    }
    return &hall, nil
}
