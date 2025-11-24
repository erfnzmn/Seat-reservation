package shows

import (
	"errors"
	"time"
)

type Service interface {
	GetAllShows() ([]Show, error)
	GetShowByID(id uint) (*Show, error)
	CreateShow(input CreateShowInput) (*Show, error)
	UpdateShow(id uint, input UpdateShowInput) (*Show, error)
	DeleteShow(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

type CreateShowInput struct {
	HallID    uint      `json:"hall_id"`
	MovieName string    `json:"movie_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type UpdateShowInput struct {
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

func (s *service) CreateShow(input CreateShowInput) (*Show, error) {

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

func (s *service) UpdateShow(id uint, input UpdateShowInput) (*Show, error) {

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

func (s *service) DeleteShow(id uint) error {
	return s.repo.Delete(id)
}
