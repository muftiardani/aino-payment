# 💳 AinoPay - Payment Management System

<div align="center">

![AinoPay](https://img.shields.io/badge/AinoPay-v1.0.0-blue)
![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)
![Nuxt](https://img.shields.io/badge/Nuxt-4.2.1-00DC82?logo=nuxt.js)
![Vue](https://img.shields.io/badge/Vue-3.5.25-4FC08D?logo=vue.js)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?logo=postgresql)
![License](https://img.shields.io/badge/license-MIT-green)

**Modern full-stack payment management application built with Go and Nuxt.js**

[Features](#-features) • [Tech Stack](#-tech-stack) • [Quick Start](#-quick-start) • [Documentation](#-documentation)

</div>

---

## 📖 Overview

**AinoPay** adalah aplikasi manajemen pembayaran full-stack yang modern dan scalable, dirancang untuk membantu pengguna melacak dan mengelola transaksi pembayaran mereka dengan mudah. Aplikasi ini dibangun dengan arsitektur yang bersih, menggunakan teknologi terkini untuk performa optimal dan pengalaman pengguna yang luar biasa.

### ✨ Highlights

- 🔐 **Secure Authentication** - JWT-based auth dengan refresh token mechanism
- 📊 **Rich Dashboard** - Real-time statistics dan interactive charts
- 🔍 **Advanced Filtering** - Multi-criteria search dan filtering
- 📱 **Responsive Design** - Mobile-first dengan dark mode support
- 🚀 **High Performance** - Optimized queries dan client-side caching
- 🎨 **Modern UI/UX** - Beautiful interface dengan smooth animations
- 📤 **Data Export** - Export payments ke CSV
- 🔄 **Real-time Updates** - Instant feedback dengan toast notifications

---

## 🚀 Tech Stack

### Backend (Go)

| Technology     | Version | Purpose                |
| -------------- | ------- | ---------------------- |
| **Go**         | 1.25    | Programming language   |
| **Gin**        | 1.11.0  | Web framework          |
| **GORM**       | 1.31.1  | ORM & database toolkit |
| **PostgreSQL** | 15+     | Relational database    |
| **JWT**        | 5.3.0   | Authentication         |
| **Bcrypt**     | -       | Password hashing       |
| **Swagger**    | 1.16.6  | API documentation      |

### Frontend (Nuxt.js)

| Technology      | Version | Purpose            |
| --------------- | ------- | ------------------ |
| **Nuxt.js**     | 4.2.1   | Vue meta-framework |
| **Vue**         | 3.5.25  | UI framework       |
| **TypeScript**  | -       | Type safety        |
| **TailwindCSS** | 6.14.0  | Styling            |
| **Pinia**       | 0.11.3  | State management   |
| **Chart.js**    | 4.5.1   | Data visualization |
| **Day.js**      | 1.11.19 | Date handling      |
| **Zod**         | 4.1.13  | Schema validation  |

---

## 🎯 Features

### 🔐 Authentication & Security

- ✅ User registration dengan email validation
- ✅ Secure login dengan JWT tokens
- ✅ Refresh token mechanism untuk session management
- ✅ Password reset flow (forgot password)
- ✅ Role-based access control (User/Admin)
- ✅ Rate limiting (10 req/s per IP)
- ✅ CORS protection
- ✅ Bcrypt password hashing

### 💰 Payment Management

- ✅ **CRUD Operations**: Create, Read, Update, Delete payments
- ✅ **Advanced Filtering**:
  - Filter by status (pending, completed, failed, refunded)
  - Search in description
  - Amount range filter (min/max)
  - Date range filter (start/end date)
- ✅ **Pagination**: Efficient data loading
- ✅ **CSV Export**: Export filtered payments
- ✅ **Categories**: Subscription, Purchase, Service, Donation, Other
- ✅ **Payment Methods**: Bank Transfer, Credit Card, E-Wallet, Cash

### 📊 Dashboard & Analytics

- ✅ **Statistics Cards**:
  - Total payments count
  - Completed payments
  - Pending payments
  - Total amount
- ✅ **Monthly Earnings Chart**: Interactive line chart dengan year selector
- ✅ **Recent Activity**: Quick view of latest 5 payments
- ✅ **Real-time Updates**: Instant data refresh

### 🎨 UI/UX Features

- ✅ **Dark Mode**: Persistent theme preference
- ✅ **Responsive Design**: Mobile, tablet, desktop optimized
- ✅ **Toast Notifications**: Success, error, warning, info messages
- ✅ **Loading States**: Skeleton loaders & spinners
- ✅ **Empty States**: Informative placeholders
- ✅ **Smooth Animations**: Fade, slide, scale transitions
- ✅ **Form Validation**: Real-time validation dengan error messages
- ✅ **Keyboard Shortcuts**: ESC to close modals

---

## 📁 Project Structure

```
ainopay-web/
├── ainopay-server/              # Backend (Go)
│   ├── cmd/server/              # Application entry point
│   ├── internal/
│   │   ├── config/              # Configuration
│   │   ├── database/            # DB connection & migrations
│   │   ├── handlers/            # HTTP handlers (controllers)
│   │   ├── middleware/          # Auth, CORS, logging, etc.
│   │   ├── models/              # Domain entities
│   │   ├── repositories/        # Data access layer
│   │   ├── services/            # Business logic
│   │   └── utils/               # Utilities (JWT, password, etc.)
│   ├── docs/                    # Swagger documentation
│   ├── go.mod                   # Go dependencies
│   └── .env                     # Environment variables
│
├── app/                         # Frontend (Nuxt.js)
│   ├── assets/                  # Static assets & styles
│   ├── components/              # Vue components
│   │   ├── common/              # Reusable UI components
│   │   ├── dashboard/           # Dashboard components
│   │   ├── icons/               # Icon components
│   │   └── layout/              # Layout components
│   ├── composables/             # Composition functions
│   ├── layouts/                 # Page layouts
│   ├── middleware/              # Route middleware
│   ├── pages/                   # File-based routing
│   ├── plugins/                 # Nuxt plugins
│   ├── stores/                  # Pinia stores
│   ├── types/                   # TypeScript definitions
│   └── utils/                   # Utility functions
│
├── public/                      # Public static files
├── nuxt.config.ts              # Nuxt configuration
├── tailwind.config.js          # Tailwind configuration
├── package.json                # Frontend dependencies
└── README.md                   # This file
```

---

## 🚀 Quick Start

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.25 or higher ([Download](https://golang.org/dl/))
- **Node.js** 18+ and npm ([Download](https://nodejs.org/))
- **PostgreSQL** 15+ ([Download](https://www.postgresql.org/download/))

### 1️⃣ Clone Repository

```bash
git clone <repository-url>
cd ainopay-web
```

### 2️⃣ Setup Database

```bash
# Login ke PostgreSQL
psql -U postgres

# Buat database
CREATE DATABASE ainopay;

# Exit
\q
```

### 3️⃣ Setup Backend

```bash
# Navigate ke backend directory
cd ainopay-server

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env dan sesuaikan dengan konfigurasi database Anda
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=your_password
# DB_NAME=ainopay
# JWT_SECRET=your-secret-key-here

# Run server
go run cmd/server/main.go
```

Backend akan berjalan di `http://localhost:8080`

### 4️⃣ Setup Frontend

```bash
# Buka terminal baru, navigate ke root directory
cd ainopay-web

# Install dependencies
npm install

# Create .env file (optional, default sudah sesuai)
echo "NUXT_PUBLIC_API_BASE=http://localhost:8080/api" > .env

# Run development server
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

### 5️⃣ Access Application

1. Buka browser dan akses `http://localhost:3000`
2. Register akun baru atau login dengan credentials yang sudah ada
3. Mulai mengelola payments Anda! 🎉

---

## 📚 Documentation

### API Documentation

Swagger UI tersedia di: `http://localhost:8080/swagger/index.html`

---

<div align="center">

**Made with ❤️ using Go and Nuxt.js**

</div>
