# Fleetify System - Technical Test

Fullstack invoice management system built with Go (Fiber) and Next.js.

## 🚀 Features
- **Zero-Trust Backend**: Total calculation is handle by the server.
- **ACID Transactions**: Ensures data integrity for invoice and details.
- **Asynchronous Webhooks**: Notifications send via Goroutines.
- **JWT Authentication**: Secure access for Admin and Kerani roles.

## 🛠 Tech Stack
- **Backend**: Go, Gorm, Fiber, PostgreSQL.
- **Infrastructure**: Docker & Docker Compose.

## 🏃 How to Run
1. **Clone the repository**
2. **Setup Environment Variables**
    - Copy `backend/.env.example` to `backend/.env`
    - Update `WEB_HOOK` URL from [webhook.site](https://webhook.site)
3. **Run with Docker**

    ```bash
   docker-compose up --build

## 🔑 Access API
- Backend: http://localhost:8080
- Mater Items (Public): GET /api/items

## 📂 Folder Structure

### Backend

```text
.
├── backend/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   ├── utils/
│   ├── .env.example
│   ├── Dockerfile
│   └── main.go
├── .gitignore
├── docker-compose.yml
└── README.md