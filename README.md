# 🧾 Job Application Tracker API

A simple RESTful API built in Go (Golang) to help you track your job applications — manage statuses, add notes, and analyze your progress. 🔍

---

## 🚀 Features

- ✅ User registration and login (JWT-based authentication)
- 📬 Add new job applications with status and notes
- 📊 View your applications by status
- 📈 Analytics: count of applications by status
- 🔐 Protected endpoints using JWT middleware

---

## 🛠 Tech Stack

- **Golang** (Gin framework)
- **PostgreSQL**
- **GORM** (ORM for Go)
- **JWT** for auth
- **dotenv** for environment variables
- (Optional) **Postman collection** for testing

---

## 📂 Project Structure

job-tracker/
│
├── cmd/server # main.go entry point
├── internal/
│ ├── database/ # DB connection
│ ├── handlers/ # HTTP handlers
│ ├── middleware/ # JWT auth middleware
│ ├── models/ # GORM models
│ └── utils/ # Utility functions (e.g., hashing, JWT generation)
├── .env # DB credentials and secret key
├── go.mod / go.sum
└── README.md

yaml
Копировать
Редактировать

---

## 🧪 API Endpoints

| Method | Endpoint                | Description                        | Auth Required |
|--------|-------------------------|------------------------------------|---------------|
| POST   | `/auth/register`        | Register a new user                | ❌            |
| POST   | `/auth/login`           | Login and get JWT token            | ❌            |
| POST   | `/applications`         | Create new application             | ✅            |
| GET    | `/applications`         | Get all user applications          | ✅            |
| PATCH  | `/applications/:id`     | Update application status          | ✅            |
| GET    | `/applications/stats`   | Get stats by application status    | ✅            |

---

## ⚙️ How to Run Locally

> Make sure you have Go and PostgreSQL installed

1. Clone the repo:

```bash
git clone https://github.com/your-username/job-application-tracker.git
cd job-application-tracker
Create .env file:

env
Копировать
Редактировать
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=job_tracker
JWT_SECRET=super-secret-key
Run the server:

bash
Копировать
Редактировать
go run cmd/server/main.go
Server will start on: http://localhost:8080

🧪 Testing with Postman
Register → /auth/register

Login → Get JWT token

Copy token to Authorization header as:
Bearer <your-token>

Access protected routes /applications, /applications/stats, etc.

📌 Future Improvements
Add email or Telegram reminders

Add pagination and search

Deploy with Docker

Build a frontend with React or Vue

🧑‍💻 Author
Made with ❤️ by @your-name
