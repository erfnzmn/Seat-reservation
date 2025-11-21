# Database Design — Cinema Seat Reservation System

## 1. Overview

The database schema is designed to support:

- Multiple users
- Multiple shows (movie + time)
- Seat layouts per show
- Safe seat reservations
- Waiting list for fully booked shows

The core tables are:

- `users`
- `shows`
- `seats`
- `reservations`
- `waiting_list`

All tables are designed for MySQL.

---

## 2. Entities and Relationships

### **Users**

- Represents customers using the system.
- A user can:
  - Make reservations
  - Join waiting lists

### **Shows**

- Represents a specific movie showtime.
- Each show:
  - Has a start time, end time, and hall name
  - Has a finite number of seats

### **Seats**

- Represents physical seats for a given show.
- Each seat:
  - Belongs to one show
  - Has row/number/label information

### **Reservations**

- Represents a confirmed or cancelled booking for a specific seat in a show.
- Each reservation:
  - Belongs to one user
  - Belongs to one show
  - Targets one seat

### **Waiting List**

- Represents users who want to attend a show that is currently fully booked.
- Each waiting entry:
  - Belongs to one user
  - Belongs to one show
  - Is processed when a seat becomes available

---

## 3. Table Definitions

### 3.1 `users` Table

Stores basic user information and credentials.
