package waitinglist

import (
	"encoding/json"
	"log"

	"seat-reservation/internals/reservation"
	"seat-reservation/pkg/rabbitmq"
)

type seatAvailableEvent struct {
	ShowID uint `json:"show_id"`
	SeatID uint `json:"seat_id"`
}

// rmq: اتصال RabbitMQ
// waitRepo: ریپوی WaitingList
// resService: سرویس Reservation
func StartWorker(rmq *rabbitmq.RabbitMQ, waitRepo Repository, resService reservation.Service) {
	if rmq == nil {
		log.Println("[WaitingWorker] RabbitMQ is nil → worker not started")
		return
	}

	msgs, err := rmq.Consume("seat.available")
	if err != nil {
		log.Println("[WaitingWorker] consume error:", err)
		return
	}

	go func() {
		log.Println("[WaitingWorker] started, listening on 'seat.available' queue")

		for msg := range msgs {
			var ev seatAvailableEvent
			if err := json.Unmarshal(msg.Body, &ev); err != nil {
				log.Println("[WaitingWorker] invalid message:", err)
				continue
			}

			log.Printf("[WaitingWorker] seat.available received: show=%d seat=%d\n", ev.ShowID, ev.SeatID)

			// ۱) نفر بعدی در صف
			w, err := waitRepo.GetNextInQueue(ev.ShowID)
			if err != nil {
				log.Println("[WaitingWorker] no waiting user or db error:", err)
				continue
			}

			// ۲) ساخت رزرو برای این نفر
			input := reservation.CreateReservationInput{
				ShowID:    ev.ShowID,
				SeatID:    ev.SeatID,
				UserName:  w.UserName,
				UserPhone: w.UserPhone,
			}

			res, err := resService.CreateReservation(input)
			if err != nil {
				log.Println("[WaitingWorker] failed to create reservation from waiting list:", err)
				continue
			}

			// ۳) آپدیت waiting_list
			if err := waitRepo.MarkAsAssigned(w.ID, ev.SeatID); err != nil {
				log.Println("[WaitingWorker] failed to mark waiting list as assigned:", err)
				continue
			}

			log.Printf("[WaitingWorker] waiting user assigned → waiting_id=%d reservation_id=%d\n", w.ID, res.ID)
		}

		log.Println("[WaitingWorker] msgs channel closed, worker exiting")
	}()
}
