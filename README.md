# Cinema Seat Reservation System

A backend system built during a mentorship challenge to manage movie halls, shows, seat maps, reservations, cancellations, and waiting lists.  
The goal was to design the entire system from scratch based on a concise project description and implement it using modern backend practices.

---

## 1. Project Overview

The API enables clients to:

- Browse halls, shows, and seat layouts  
- Reserve seats for upcoming screenings  
- Cancel reservations  
- Join a waiting list when shows are sold out  
- Automatically receive a seat when one becomes available  

Internally, the project follows a modular, layered architecture with clean separation between HTTP handlers, business logic, database persistence, and infrastructure.

---

## 2. System Architecture

### Technology Stack

- **Go + Echo** for API routing  
- **MySQL** for database persistence  
- **Redis** for caching and distributed seat-locks  
- **RabbitMQ** for asynchronous waiting-list processing  
- **Docker + docker-compose** for local orchestration  

### Structure

Each domain (halls, shows, seats, reservations, waiting list) includes:

- `handler.go` – HTTP layer  
- `service.go` – business logic  
- `repository.go` – database operations  
- `model.go` – data structures  

The server loads configs, initializes infrastructure, registers routes, and starts the HTTP engine.

---

## 3. Core Features

- Seat reservation with concurrency-safe locking via Redis  
- Reservation cancellation that triggers RabbitMQ events  
- Automatic waiting-list assignment when seats free up  
- Show and hall management  
- Automatic database seeding on first run  

---

## 4. Database Schema

SQL migrations define the schema for:

- Halls  
- Shows  
- Seats  
- Reservations  
- Waiting list  

Stored in the `migrations/` directory and applied automatically at startup.

---

## 5. Containers & Local Development

**Dockerfile** builds a lightweight Go binary using multi-stage builds.  
**docker-compose.yml** orchestrates:

- App container  
- MySQL  
- Redis  
- RabbitMQ  
- Internal network + persistent volumes  

Start everything with:

```bash
docker compose up --build

---

Repository Structure

cmd/server/         → Entry point & server bootstrap
internals/          → Domain modules (halls, shows, seats, reservations, waitinglist)
pkg/                → Infrastructure utilities (Redis, RabbitMQ, middleware)
configs/            → YAML configuration files
docs/               → Architecture notes
migrations/         → SQL schema definitions
Makefile            → Developer commands
Dockerfile          → Build instructions
docker-compose.yml  → Local environment orchestration

