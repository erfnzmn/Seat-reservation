package seeder

import (
    "log"

    "gorm.io/gorm"

    "seat-reservation/internals/halls"
    "seat-reservation/internals/seats"
	
)

// اجرای فقط یک‌بار
func SeedInitialData(db *gorm.DB) {
    var count int64
    db.Model(&halls.Hall{}).Count(&count)

    if count > 0 {
        log.Println("Seed already exists — skipping seeding.")
        return
    }

    log.Println("Seeding initial halls and seats...")

    // ساخت سالن‌ها
    hallData := []struct {
        Name       string
        TotalSeats int
    }{
        {"Hall 1", 120},
        {"Hall 2", 120},
        {"Hall 3", 60},
        {"Hall 4", 60},
    }

    createdHalls := make([]halls.Hall, 0)

    for _, h := range hallData {
        hall := halls.Hall{
            Name:       h.Name,
            TotalSeats: h.TotalSeats,
        }
        db.Create(&hall)
        createdHalls = append(createdHalls, hall)
    }

    // ساخت صندلی‌ها
    for _, hall := range createdHalls {
        for i := 1; i <= hall.TotalSeats; i++ {
            seat := seats.Seat{
                HallID:     hall.ID,
                SeatNumber: i,
            }
            db.Create(&seat)
        }
        log.Printf("Created %d seats for %s", hall.TotalSeats, hall.Name)
    }

    log.Println("Seeding complete ✔")
}
