# 🎬 Cinema Seat Reservation System — Overview

## 1. Introduction

The Cinema Seat Reservation System is a backend service designed to handle real-time reservation of seats across multiple cinema halls. Customers should be able to browse upcoming shows, view available seats, reserve seats, cancel reservations, and join a waiting list when all seats are sold out.  
This system aims to provide reliable, concurrent, and scalable seat management for high-traffic cinema environments.

---

## 2. Goals of the System

- Provide a clean and reliable backend API for managing movie shows, seats, and reservations.
- Ensure safe and atomic seat reservations even under heavy concurrency.
- Automatically handle cancellations and seat reassignment using a waiting list.
- Expose efficient APIs that the cinema’s mobile app or website can consume.
- Maintain high system availability using caching and event-driven components.

---

## 3. Actors

### **Customer (User)**

- Browses shows
- Views seat maps
- Reserves seats
- Cancels reservations
- Joins waiting list if seats are full

### **Staff/Admin**

- Creates shows
- Manages halls and seat layouts

### **System Services**

- Redis cache (seat map caching)
- RabbitMQ (waiting list processing and event handling)

---

## 4. High-Level Features

- User registration & authentication (JWT-based).
- Show management (create, list, details).
- Seat management (layout, availability, seat map caching).
- Real-time seat reservation with concurrency controls.
- Reservation cancellation + auto seat reassignment.
- Waiting list mechanism for fully booked shows.
- Event-driven workflows using RabbitMQ.
- Caching of seat maps and availability using Redis.

---

## 5. Core Use Cases

### **UC1 — Browse Shows**

The user retrieves all upcoming shows and selects one.

### **UC2 — View Seat Map**

User views the hall layout and sees each seat’s status.

### **UC3 — Reserve Seat(s)**

- User selects a seat or group of seats.
- System validates availability.
- Reservation is confirmed atomically.
- If the show is full → user is added to waiting list.

### **UC4 — Cancel Reservation**

- User cancels before the show starts.
- Seat becomes available.
- Event is published to RabbitMQ.

### **UC5 — Automatic Waiting List Processing**

Triggered by the cancellation event:

- Pop the next user from waiting list.
- Automatically assign the newly freed seat.
- Notify user (future enhancement).

---

## 6. System Architecture Overview

### **Key Architectural Components**

#### **1. Echo API Layer**

- Handles HTTP requests and responses  
- Performs request validation  
- Connects the client to the service layer  
- Implements routing and middleware (e.g., authentication)

#### **2. Services Layer**

- Contains all business logic  
- Orchestrates reservation creation, cancellation, and seat state transitions  
- Communicates with repositories, Redis, and RabbitMQ  
- Ensures transactional integrity and concurrency safety  

#### **3. Repository Layer (MySQL)**

- Responsible for persistent storage  
- Manages SQL queries and transactions  
- Ensures ACID guarantees during seat reservation and cancellation  

#### **4. Redis Cache**

- Stores seat maps for fast access  
- Reduces database load during high-demand showtimes  
- Seat map is invalidated/updated whenever reservation or cancellation occurs  

#### **5. RabbitMQ**

- Used for event-driven workflows  
- Handles asynchronous processes such as processing the waiting list  
- Ensures the reservation logic remains responsive even under heavy load  

#### **6. Waiting List Worker**

- Listens to “seat.available” events  
- Fetches the next user from the waiting list  
- Automatically assigns the freed seat  
- Prevents race conditions during automatic reassignment  

---

## 7. Key Technical Challenges

### **1. Concurrency control in seat reservation**

Multiple users may attempt to reserve the same seat simultaneously.  
The system must guarantee **no overbooking** using:

- SQL transactions  
- Optimistic or pessimistic locking  
- Optional Redis-based distributed locks  
- Event-driven updates to avoid race conditions  

### **2. High-demand showtimes**

Popular movies generate sudden traffic spikes.  
Challenges include:

- Thousands of seat map requests  
- Many concurrent reservation attempts  
- Potential DB overload  

Solutions:

- Redis caching  
- Asynchronous task processing  
- Efficient, read-optimized API endpoints  

### **3. Waiting List Mechanism**

When a show is fully booked:

- Users are added to a waiting list  
- A “seat.available” event is triggered on cancellation  
- The worker assigns seats to users automatically  

### **4. Real-time seat map updates**

Seat availability must always be accurate.

Techniques:

- Cache seat map in Redis  
- Update/invalidate cache whenever seat state changes  
- Avoid heavy SQL queries during peak hours  

---

## 8. Modules Breakdown

### **/internals/shows**

- Show model  
- Create/list/get showtime information  
- Validate showtime availability  

### **/internals/seats**

- Seat model and hall layout  
- Seat availability status  
- Integration with Redis seat map cache  

### **/internals/reservation**

- Reserve seats with concurrency safety  
- Cancel reservations  
- Manage waiting list (`waitings.go`)  
- Publish events to RabbitMQ  

### **/internals/users**

- User registration and login  
- JWT-based authentication  
- Middleware for protected routes  

---

## 9. Reservation Flow Summary

### **Step 1 — User selects a show**

The backend returns showtime details and seat map.

### **Step 2 — User selects seat(s)**

The system checks availability using both DB and Redis.

### **Step 3 — Reservation logic**

- If available → confirm and save  
- If full → add user to waiting list  

### **Step 4 — Cancellation**

- Seat becomes free  
- “seat.available” event is published  

### **Step 5 — Automatic waiting list processing**

- Worker receives the event  
- Pops the next user from waiting list  
- Automatically creates a reservation  

---

## 10. Future Enhancements

- Payment integration  
- Notification service (SMS/Email)  
- Group-seat recommendation system  
- Admin dashboard  
- Show-time analytics and reporting  

---

## 11. Project Status

This document defines the high-level structure and scope of the Cinema Seat Reservation System.  
All future development steps—database design, API development, caching strategy, and event-driven mechanisms—will follow the architecture described above.
