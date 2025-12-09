# PostgreSQL Database Setup Guide

## Langkah 1: Install PostgreSQL

### Windows

1. **Download PostgreSQL**
   - Kunjungi: https://www.postgresql.org/download/windows/
   - Download installer (PostgreSQL 15 atau 16)

2. **Install PostgreSQL**
   - Jalankan installer
   - Port default: 5432
   - Set password untuk user `postgres` (ingat password ini!)
   - Install semua components (termasuk pgAdmin)

3. **Verify Installation**

   ```powershell
   # Check PostgreSQL version
   psql --version

   # Check if service running
   Get-Service postgresql*
   ```

## Langkah 2: Create Database

### Menggunakan psql (Command Line)

```powershell
# Login ke PostgreSQL
psql -U postgres

# Di psql prompt, create database
CREATE DATABASE ainopay;

# Verify database created
\l

# Connect to database
\c ainopay

# Exit psql
\q
```

### Menggunakan pgAdmin (GUI)

1. Buka pgAdmin
2. Connect ke PostgreSQL server
3. Right-click "Databases" → Create → Database
4. Database name: `ainopay`
5. Owner: `postgres`
6. Click Save

## Langkah 3: Configure Backend

File `.env` sudah ada di `ainopay-server/.env.example`. Copy dan sesuaikan:

```bash
# Di ainopay-server directory
cp .env.example .env
```

Edit `.env` dan sesuaikan dengan konfigurasi PostgreSQL Anda:

```env
# Server Configuration
PORT=8080
GIN_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=YOUR_POSTGRES_PASSWORD_HERE  # Ganti dengan password Anda
DB_NAME=ainopay
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=ainopay-secret-key-2025
JWT_EXPIRATION=24h

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000
```

## Langkah 4: Test Database Connection

```powershell
# Di ainopay-server directory
cd ainopay-server

# Run server (akan auto-migrate dan seed data)
go run cmd/server/main.go
```

Expected output:

```
Database connected successfully
Running database migrations...
Database migration completed successfully
Seeding database...
Database seeding completed successfully
Server starting on port 8080
```

## Langkah 5: Verify Database Tables

### Menggunakan psql

```powershell
# Login ke database
psql -U postgres -d ainopay

# List tables
\dt

# Check users table
SELECT * FROM users;

# Check categories
SELECT * FROM categories;

# Check payment_methods
SELECT * FROM payment_methods;

# Exit
\q
```

Expected tables:

- `users`
- `payments`
- `categories`
- `payment_methods`

### Menggunakan pgAdmin

1. Buka pgAdmin
2. Navigate: Servers → PostgreSQL → Databases → ainopay → Schemas → public → Tables
3. Verify 4 tables ada

## Langkah 6: Test API Endpoints

### Health Check

```powershell
curl http://localhost:8080/health
```

Expected: `{"status":"ok"}`

### Register User

```powershell
curl -X POST http://localhost:8080/api/auth/register `
  -H "Content-Type: application/json" `
  -d '{
    "email": "admin@ainopay.com",
    "password": "admin123",
    "full_name": "Admin User"
  }'
```

### Login

```powershell
curl -X POST http://localhost:8080/api/auth/login `
  -H "Content-Type: application/json" `
  -d '{
    "email": "admin@ainopay.com",
    "password": "admin123"
  }'
```

Save token dari response!

### Get Categories (dengan token)

```powershell
curl -X GET http://localhost:8080/api/categories `
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Troubleshooting

### Error: "password authentication failed"

**Solution**:

- Check password di `.env` sudah benar
- Reset password PostgreSQL jika lupa

### Error: "database ainopay does not exist"

**Solution**:

```powershell
psql -U postgres
CREATE DATABASE ainopay;
\q
```

### Error: "could not connect to server"

**Solution**:

- Check PostgreSQL service running
- Verify port 5432 tidak digunakan aplikasi lain
- Check firewall settings

### Error: "relation does not exist"

**Solution**:

- Restart server untuk run migrations
- Check migration logs

## Database Schema

Setelah migration, struktur database:

### users

```sql
id              UUID PRIMARY KEY
email           VARCHAR UNIQUE NOT NULL
password_hash   VARCHAR NOT NULL
full_name       VARCHAR NOT NULL
role            VARCHAR(20) DEFAULT 'user'
created_at      TIMESTAMP
updated_at      TIMESTAMP
```

### categories

```sql
id              UUID PRIMARY KEY
name            VARCHAR NOT NULL
description     TEXT
created_at      TIMESTAMP
updated_at      TIMESTAMP
```

### payment_methods

```sql
id              UUID PRIMARY KEY
name            VARCHAR NOT NULL
code            VARCHAR UNIQUE NOT NULL
is_active       BOOLEAN DEFAULT true
created_at      TIMESTAMP
updated_at      TIMESTAMP
```

### payments

```sql
id                  UUID PRIMARY KEY
user_id             UUID NOT NULL (FK -> users)
amount              DECIMAL(15,2) NOT NULL
status              VARCHAR(20) DEFAULT 'pending'
payment_method_id   UUID NOT NULL (FK -> payment_methods)
category_id         UUID NOT NULL (FK -> categories)
description         TEXT
transaction_date    TIMESTAMP NOT NULL
created_at          TIMESTAMP
updated_at          TIMESTAMP
```

## Seed Data

Otomatis di-seed saat server start:

**Categories:**

- Subscription
- Purchase
- Service
- Donation
- Other

**Payment Methods:**

- Bank Transfer
- Credit Card
- E-Wallet
- Cash

## Next Steps

1. ✅ PostgreSQL installed
2. ✅ Database `ainopay` created
3. ✅ `.env` configured
4. ✅ Server running dan migrations success
5. ✅ Test API endpoints
6. 🔄 Run frontend: `npm run dev`
7. 🔄 Test full application flow

## Quick Reference

```powershell
# Start PostgreSQL (if not auto-start)
net start postgresql-x64-15  # Sesuaikan dengan versi Anda

# Stop PostgreSQL
net stop postgresql-x64-15

# Check PostgreSQL status
Get-Service postgresql*

# Login to database
psql -U postgres -d ainopay

# Backup database
pg_dump -U postgres ainopay > backup.sql

# Restore database
psql -U postgres ainopay < backup.sql
```
