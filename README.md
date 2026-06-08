# 🚀 SIMPRO-BE

> Backend API for **Simple Inventory Management & Procurement System (SIMPRO)** built with **Golang**, **Gin**, **MySQL**, and **JWT Authentication**.

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00ADD8)
![MySQL](https://img.shields.io/badge/MySQL-Database-4479A1?logo=mysql)
![JWT](https://img.shields.io/badge/JWT-Authentication-orange)

---

## 📖 Overview

SIMPRO-BE is a backend application designed to manage inventory and procurement processes across multiple branches.

### ✨ Current Features

* 🔐 User Authentication
* 🔑 JWT Authorization
* 🛡️ Protected API Routes
* 🔒 Password Hashing with Bcrypt
* 🗄️ MySQL Integration
* ⚡ Layered Architecture (Handler → Service → Repository)

---

## 🏗️ Architecture

```text
Client
   │
   ▼
Routes
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
MySQL Database
```

---

## 🛠️ Tech Stack

| Technology | Description         |
| ---------- | ------------------- |
| 🐹 Golang  | Backend Language    |
| 🌐 Gin     | HTTP Framework      |
| 🗄️ MySQL  | Database            |
| 📦 GORM    | ORM                 |
| 🔑 JWT     | Authentication      |
| 🔒 Bcrypt  | Password Encryption |

---

## 📂 Project Structure

```text
SIMPRO-BE
│
├── cmd
│   └── server
│       └── main.go
│
├── config
│   ├── config.go
│   └── database.go
│
├── internal
│   ├── dto
│   │   └── auth
│   │
│   ├── handler
│   ├── middleware
│   ├── model
│   ├── repository
│   ├── service
│   └── utils
│
├── routes
│
├── schema.sql
├── .env.example
└── README.md
```

---

## ⚙️ Getting Started

### 1️⃣ Clone Repository

```bash
git clone https://github.com/KuroKitsune9/SIMPRO-BE.git
cd SIMPRO-BE
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Configure Environment

Create `.env` file:

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_NAME=simpro
DB_USER=root
DB_PASS=

JWT_SECRET=simpro-secret-key
```

### 4️⃣ Create Database

```sql
CREATE DATABASE simpro;
```

### 5️⃣ Import Schema

```bash
mysql -u root -p simpro < schema.sql
```

### 6️⃣ Run Application

```bash
go run cmd/server/main.go
```

Application will be available at:

```text
http://localhost:8080
```

---

## 🔐 Authentication

### 📝 Register

**POST** `/api/auth/register`

Request:

```json
{
  "name": "John Doe",
  "username": "john",
  "password": "123456",
  "role": "InventoryControl",
  "branchCode": "001"
}
```

---

### 🔑 Login

**POST** `/api/auth/login`

Request:

```json
{
  "username": "john",
  "password": "123456"
}
```

Response:

```json
{
  "success": true,
  "token": "JWT_TOKEN"
}
```

---

### 👤 Profile

**GET** `/api/auth/profile`

Header:

```http
Authorization: Bearer JWT_TOKEN
```

Response:

```json
{
  "success": true,
  "userID": 1,
  "username": "john"
}
```

---

## 🔄 Authentication Flow

```text
📝 Register
      │
      ▼
🔒 Password Hashing (Bcrypt)
      │
      ▼
💾 Save User
      │
      ▼
🔑 Login
      │
      ▼
🎟️ Generate JWT
      │
      ▼
📦 Return Token
      │
      ▼
🛡️ JWT Middleware
      │
      ▼
✅ Access Protected Endpoint
```

---

## 📊 Project Status

### ✅ Completed

* [x] Project Setup
* [x] MySQL Connection
* [x] User Registration
* [x] User Login
* [x] JWT Generation
* [x] JWT Middleware
* [x] Protected Route

### 🚧 In Progress

* [ ] Branch Management
* [ ] Item Management
* [ ] Inventory Management

### 🎯 Planned

* [ ] Purchase Order
* [ ] Goods Receipt
* [ ] Goods Return
* [ ] Dashboard API
* [ ] Reporting
* [ ] Audit Log

---

## 👥 Contributors

| Name             | Role              |
| ---------------- | ----------------- |
| Muharafi Dalilah | Backend Developer |

---

## 📄 License

This project is developed for educational and academic purposes.

⭐ If you find this project useful, consider giving it a star.

