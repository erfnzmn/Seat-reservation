package reservation

import (
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"context"
)

var (
	ErrSeatAlreadyReserved = errors.New("seat already reserved for this show")
	ErrSeatLockFailed      = errors.New("could not lock seat for reservation")
)

type Service interface {
	CreateReservation(input CreateReservationInput) (*Reservation, error)
	CancelReservation(id uint) error
	GetReservation(id uint) (*Reservation, error)
}

type service struct {
	repo  Repository
	redis *redis.Client
}

func NewService(repo Repository, redis *redis.Client) Service {
	return &service{repo: repo, redis: redis}
}


func (s *service) CreateReservation(input CreateReservationInput) (*Reservation, error) {
	ctx := context.Background()

	// 1) Build Redis key
	lockKey := fmt.Sprintf("lock:seat:%d:show:%d", input.SeatID, input.ShowID)

	// 2) Try Lock (3 seconds)
	ok, err := s.redis.SetNX(ctx, lockKey, 1, 3*time.Second).Result()
	if err != nil || !ok {
		return nil, ErrSeatLockFailed
	}
	defer s.redis.Del(ctx, lockKey)

	// 3) Check double booking
	existing, err := s.repo.GetBySeatAndShow(input.SeatID, input.ShowID)
	if err == nil && existing != nil && existing.Status == "confirmed" {
		return nil, ErrSeatAlreadyReserved
	}

	res := &Reservation{
		ShowID:    input.ShowID,
		SeatID:    input.SeatID,
		UserName:  input.UserName,
		UserPhone: input.UserPhone,
		Status:    "confirmed",
	}

	if err := s.repo.Create(res); err != nil {
		return nil, err
	}

	return res, nil
}



func (s *service) CancelReservation(id uint) error {
	return s.repo.Cancel(id)
}


func (s *service) GetReservation(id uint) (*Reservation, error) {
	return s.repo.GetByID(id)
}
