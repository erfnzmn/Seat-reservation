package reservation

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"seat-reservation/pkg/rabbitmq"
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
	repo   Repository
	redis  *redis.Client
	rabbit *rabbitmq.RabbitMQ
}

func NewService(repo Repository, redis *redis.Client, rabbit *rabbitmq.RabbitMQ) Service {
	return &service{
		repo:   repo,
		redis:  redis,
		rabbit: rabbit,
	}
}

func (s *service) CreateReservation(input CreateReservationInput) (*Reservation, error) {
	ctx := context.Background()

	lockKey := fmt.Sprintf("lock:seat:%d:show:%d", input.SeatID, input.ShowID)

	ok, err := s.redis.SetNX(ctx, lockKey, 1, 3*time.Second).Result()
	if err != nil || !ok {
		return nil, ErrSeatLockFailed
	}
	defer s.redis.Del(ctx, lockKey)

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
    // 1) رزرو را پیدا کن تا show_id و seat_id را داشته باشیم
    res, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }

    // 2) منطق لغو را به ریپازیتوری بسپار (متد Cancel که قبلاً داشتی)
    if err := s.repo.Cancel(id); err != nil {
        return err
    }

    // 3) انتشار رویداد seat.available برای Worker
    type SeatAvailableEvent struct {
        ShowID uint `json:"show_id"`
        SeatID uint `json:"seat_id"`
    }

    event := SeatAvailableEvent{
        ShowID: res.ShowID,
        SeatID: res.SeatID,
    }

    if err := s.rabbit.Publish("seat.available", event); err != nil {
        log.Println("[Reservation] Failed to publish seat.available:", err)
    } else {
        log.Printf("[Reservation] Published seat.available → show=%d seat=%d\n",
            res.ShowID, res.SeatID)
    }

    return nil
}


func (s *service) GetReservation(id uint) (*Reservation, error) {
	return s.repo.GetByID(id)
}
