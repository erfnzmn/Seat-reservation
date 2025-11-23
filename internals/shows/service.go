package shows

import (
	"time"
	"errors"
)

type Service interface {
	GetAllShows() ([]Show, error)
	GetShowByID(id uint) (*Show, error)
	CreateShow(input CreateShowInput, adminKey string) (*Show, error)
	UpdateShow(id uint, input UpdateShowInput, adminKey string) (*Show, error)
	DeleteShow(id uint, adminKey string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

type CreateShowInput struct {
	AdminKey  string    `json:"admin_key"`
	HallID    uint      `json:"hall_id"`
	MovieName string    `json:"movie_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type UpdateShowInput struct {
	AdminKey  string    `json:"admin_key"`
	MovieName string    `json:"movie_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (s *service) GetAllShows() ([]Show, error) {
	return s.repo.GetAll()
}

func (s *service) GetShowByID(id uint) (*Show, error) {
	return s.repo.GetByID(id)
}

func (s *service) CreateShow(input CreateShowInput, adminKey string) (*Show, error) {

	if input.AdminKey != adminKey {
        return nil, errors.New("forbidden: invalid admin key")
    }

	if input.MovieName == "" {
		return nil, errors.New("movie name is required")
	}

	if input.StartTime.After(input.EndTime) {
		return nil, errors.New("start time must be before end time")
	}

	show := &Show{
		HallID:    input.HallID,
		MovieName: input.MovieName,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
	}

	if err := s.repo.Create(show); err != nil {
		return nil, err
	}

	return show, nil
}

func (s *service) UpdateShow(id uint, input UpdateShowInput, adminKey string) (*Show, error) {

    if input.AdminKey != adminKey {
        return nil, errors.New("forbidden: invalid admin key")
    }
	show, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.MovieName != "" {
		show.MovieName = input.MovieName
	}
	if !input.StartTime.IsZero() {
		show.StartTime = input.StartTime
	}
	if !input.EndTime.IsZero() {
		show.EndTime = input.EndTime
	}

	if err := s.repo.Update(show); err != nil {
		return nil, err
	}

	return show, nil
}

func (s *service) DeleteShow(id uint, adminKey string) error {

	if adminKey == "" {
        return errors.New("forbidden: invalid admin key")
    }
	return s.repo.Delete(id)
}