# AinoPay Frontend

Frontend aplikasi manajemen pembayaran menggunakan Nuxt.js 4, Vue 3, TypeScript, dan TailwindCSS.

## 🚀 Tech Stack

- **Nuxt.js** 4.2.1 - Meta-framework Vue.js
- **Vue** 3.5.25 - Progressive JavaScript framework
- **TypeScript** - Type-safe development
- **TailwindCSS** - Utility-first CSS framework
- **Pinia** - State management
- **VueUse** - Composition utilities
- **Day.js** - Date formatting

## 📁 Project Structure

```
app/
├── assets/css/main.css           # Global styles + Tailwind
├── components/                   # Vue components
│   ├── common/                   # Reusable UI components
│   └── layout/                   # Layout components
├── composables/                  # Composition functions
├── layouts/                      # Page layouts
├── middleware/                   # Route middleware
├── pages/                        # File-based routing
├── plugins/                      # Nuxt plugins
├── stores/                       # Pinia stores
└── types/                        # TypeScript types
```

## 🔧 Setup

```bash
# Install dependencies
npm install

# Configure environment (.env)
NUXT_PUBLIC_API_BASE=http://localhost:8080/api

# Run development server
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

## 🎯 Features

- ✅ Authentication (Login/Register)
- ✅ Dashboard with statistics
- ✅ Payments list with pagination
- ✅ Categories view
- ✅ Dark mode support
- ✅ Responsive design
- ✅ Toast notifications

## 🔌 Integration with Backend

Pastikan backend server running:

```bash
# Terminal 1: Backend
cd ainopay-server
go run cmd/server/main.go

# Terminal 2: Frontend
npm run dev
```

Lihat dokumentasi lengkap di [ainopay-server/README.md](ainopay-server/README.md)
