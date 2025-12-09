# AinoPay Backend Server

Backend API untuk aplikasi manajemen pembayaran menggunakan Golang, Gin Framework, dan PostgreSQL.

## 🚀 Tech Stack

- **Go** 1.25
- **Gin** - Web framework
- **GORM** - ORM untuk database
- **PostgreSQL** - Database
- **JWT** - Authentication
- **bcrypt** - Password hashing

## 📁 Project Structure

```
ainopay-server/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── config/                  # Configuration
│   ├── database/                # Database connection & migrations
│   ├── models/                  # Data models
│   ├── handlers/                # HTTP handlers (controllers)
│   ├── services/                # Business logic
│   ├── repositories/            # Data access layer
│   ├── middleware/              # Middleware (auth, cors, logger)
│   └── utils/                   # Utilities (jwt, password, response)
├── .env                         # Environment variables
├── .env.example                 # Environment template
├── go.mod                       # Go dependencies
└── Makefile                     # Development commands
```

## 🔧 Setup

### Prerequisites

- Go 1.25 or higher
- PostgreSQL 15 or higher

### Installation

1. **Clone repository** (sudah ada)

2. **Install dependencies**

   ```bash
   cd ainopay-server
   go mod download
   ```

3. **Setup PostgreSQL Database**

   ```bash
   # Login ke PostgreSQL
   psql -U postgres

   # Buat database
   CREATE DATABASE ainopay;

   # Exit
   \q
   ```

4. **Configure environment**

   File `.env` sudah dibuat dengan konfigurasi default:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=ainopay
   ```

   Sesuaikan dengan konfigurasi PostgreSQL Anda jika berbeda.

5. **Run server**

   ```bash
   # Menggunakan Makefile
   make run

   # Atau langsung dengan go
   go run cmd/server/main.go
   ```

   Server akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### Authentication

| Method | Endpoint             | Description        | Auth |
| ------ | -------------------- | ------------------ | ---- |
| POST   | `/api/auth/register` | Register user baru | ❌   |
| POST   | `/api/auth/login`    | Login user         | ❌   |
| GET    | `/api/auth/me`       | Get current user   | ✅   |

### Payments

| Method | Endpoint            | Description         | Auth |
| ------ | ------------------- | ------------------- | ---- |
| GET    | `/api/payments`     | List semua payments | ✅   |
| GET    | `/api/payments/:id` | Get payment by ID   | ✅   |
| POST   | `/api/payments`     | Create payment baru | ✅   |
| PUT    | `/api/payments/:id` | Update payment      | ✅   |
| DELETE | `/api/payments/:id` | Delete payment      | ✅   |

**Query Parameters untuk GET /api/payments:**

- `page` - Page number (default: 1)
- `limit` - Items per page (default: 10)
- `status` - Filter by status (pending, completed, failed, refunded)
- `search` - Search in description

### Categories

| Method | Endpoint          | Description           | Auth |
| ------ | ----------------- | --------------------- | ---- |
| GET    | `/api/categories` | List semua categories | ✅   |

### Payment Methods

| Method | Endpoint               | Description          | Auth |
| ------ | ---------------------- | -------------------- | ---- |
| GET    | `/api/payment-methods` | List payment methods | ✅   |

### Dashboard

| Method | Endpoint                | Description         | Auth |
| ------ | ----------------------- | ------------------- | ---- |
| GET    | `/api/dashboard/stats`  | Get statistics      | ✅   |
| GET    | `/api/dashboard/recent` | Get recent payments | ✅   |

## 🔐 Authentication

API menggunakan JWT (JSON Web Token) untuk authentication.

### Register

```bash
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "full_name": "John Doe"
}
```

### Login

```bash
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "full_name": "John Doe",
      "role": "user"
    }
  }
}
```

### Using Token

Untuk endpoint yang memerlukan authentication, sertakan token di header:

```bash
Authorization: Bearer <your-token>
```

## 💾 Database Schema

### Users

- id (UUID)
- email (unique)
- password_hash
- full_name
- role (admin/user)
- created_at, updated_at

### Payments

- id (UUID)
- user_id (FK)
- amount
- status (pending/completed/failed/refunded)
- payment_method_id (FK)
- category_id (FK)
- description
- transaction_date
- created_at, updated_at

### Categories

- id (UUID)
- name
- description
- created_at, updated_at

### Payment Methods

- id (UUID)
- name
- code
- is_active
- created_at, updated_at

## 🛠️ Development

### Available Commands

```bash
# Run server
make run

# Build binary
make build

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean

# Install dependencies
make deps
```

### Database Migrations

Migrations berjalan otomatis saat server start. GORM AutoMigrate akan membuat/update tabel sesuai dengan models.

### Seed Data

Seed data otomatis dijalankan saat server start, meliputi:

- Payment Methods (Bank Transfer, Credit Card, E-Wallet, Cash)
- Categories (Subscription, Purchase, Service, Donation, Other)

## 🧪 Testing API

### Menggunakan cURL

```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","full_name":"Test User"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Get payments (dengan token)
curl -X GET http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Menggunakan Postman/Thunder Client

1. Import collection atau buat request manual
2. Set base URL: `http://localhost:8080`
3. Untuk protected endpoints, tambahkan header:
   - Key: `Authorization`
   - Value: `Bearer YOUR_TOKEN`

## 📝 Environment Variables

| Variable        | Description              | Default               |
| --------------- | ------------------------ | --------------------- |
| PORT            | Server port              | 8080                  |
| GIN_MODE        | Gin mode (debug/release) | debug                 |
| DB_HOST         | PostgreSQL host          | localhost             |
| DB_PORT         | PostgreSQL port          | 5432                  |
| DB_USER         | Database user            | postgres              |
| DB_PASSWORD     | Database password        | postgres              |
| DB_NAME         | Database name            | ainopay               |
| DB_SSLMODE      | SSL mode                 | disable               |
| JWT_SECRET      | JWT secret key           | (set in .env)         |
| JWT_EXPIRATION  | Token expiration         | 24h                   |
| ALLOWED_ORIGINS | CORS allowed origins     | http://localhost:3000 |

## 🚨 Troubleshooting

### Database Connection Error

```
Failed to connect to database
```

**Solution**:

- Pastikan PostgreSQL running
- Check credentials di `.env`
- Pastikan database `ainopay` sudah dibuat

### Port Already in Use

```
Failed to start server: listen tcp :8080: bind: address already in use
```

**Solution**:

- Ubah PORT di `.env`
- Atau kill process yang menggunakan port 8080

### Module Not Found

```
package xxx is not in GOROOT
```

**Solution**:

```bash
go mod download
go mod tidy
```

## 📚 Next Steps

- [ ] Setup frontend Nuxt.js
- [ ] Integrate frontend dengan backend API
- [ ] Add more tests
- [ ] Add API documentation (Swagger)
- [ ] Setup Docker
- [ ] Add logging to file
- [ ] Add rate limiting
- [ ] Add request validation middleware

## 🤝 Contributing

Backend sudah siap digunakan! Selanjutnya kita akan develop frontend dengan Nuxt.js.
