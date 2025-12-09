# Backend API Testing Guide

## Prerequisites

Sebelum testing, pastikan:

1. ✅ PostgreSQL sudah terinstall dan running
2. ✅ Database `ainopay` sudah dibuat
3. ✅ File `.env` sudah dikonfigurasi dengan benar

## Setup Database

```bash
# Login ke PostgreSQL
psql -U postgres

# Buat database
CREATE DATABASE ainopay;

# Verify
\l

# Exit
\q
```

## Start Server

```bash
cd ainopay-server
go run cmd/server/main.go
```

Server akan:

1. Connect ke PostgreSQL
2. Run migrations (create tables)
3. Seed initial data
4. Start di port 8080

Expected output:

```
Database connected successfully
Running database migrations...
Database migration completed successfully
Seeding database...
Database seeding completed successfully
Server starting on port 8080
```

## Test Endpoints

### 1. Health Check

```bash
curl http://localhost:8080/health
```

Expected: `{"status":"ok"}`

### 2. Register User

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@ainopay.com",
    "password": "admin123",
    "full_name": "Admin User"
  }'
```

Expected: Success response dengan token

### 3. Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@ainopay.com",
    "password": "admin123"
  }'
```

Save the token dari response!

### 4. Get Current User

```bash
curl -X GET http://localhost:8080/api/auth/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 5. Get Categories

```bash
curl -X GET http://localhost:8080/api/categories \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Expected: List of 5 categories (Subscription, Purchase, Service, Donation, Other)

### 6. Get Payment Methods

```bash
curl -X GET http://localhost:8080/api/payment-methods \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

Expected: List of 4 payment methods (Bank Transfer, Credit Card, E-Wallet, Cash)

### 7. Create Payment

```bash
curl -X POST http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 100000,
    "payment_method_id": "GET_FROM_PAYMENT_METHODS",
    "category_id": "GET_FROM_CATEGORIES",
    "description": "Test payment",
    "transaction_date": "2025-12-09T10:00:00Z"
  }'
```

### 8. Get All Payments

```bash
curl -X GET "http://localhost:8080/api/payments?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 9. Get Dashboard Stats

```bash
curl -X GET http://localhost:8080/api/dashboard/stats \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 10. Get Recent Payments

```bash
curl -X GET http://localhost:8080/api/dashboard/recent \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Using Postman/Thunder Client

1. **Import Collection** atau create requests manually
2. **Set Environment Variables:**

   - `base_url`: `http://localhost:8080`
   - `token`: (akan diisi setelah login)

3. **Test Flow:**
   - Register → Save token
   - Login → Update token
   - Get Categories → Note category IDs
   - Get Payment Methods → Note method IDs
   - Create Payment → Use IDs from above
   - Get Payments → Verify created payment
   - Get Dashboard Stats → See statistics

## Common Issues

### Database Connection Failed

```
Failed to connect to database
```

**Fix**:

- Check PostgreSQL is running: `pg_isready`
- Verify credentials in `.env`
- Ensure database exists: `psql -U postgres -l`

### Port Already in Use

```
bind: address already in use
```

**Fix**: Change PORT in `.env` or kill process using port 8080

### Unauthorized Error

```
{"success":false,"error":"Authorization header required"}
```

**Fix**: Include `Authorization: Bearer <token>` header

## Next Steps

Setelah backend berjalan dengan baik:

1. ✅ Test semua endpoints
2. ✅ Verify data di PostgreSQL
3. 🔄 Mulai develop frontend Nuxt.js
4. 🔄 Integrate frontend dengan backend API
