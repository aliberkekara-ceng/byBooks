# 📚 byBooks Backend - Library & URL Processor API

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![Framework](https://img.shields.io/badge/Framework-Gin-008080?style=for-the-badge)](https://gin-gonic.github.io/gin/)
[![ORM](https://img.shields.io/badge/ORM-GORM-blueviolet?style=for-the-badge)](https://gorm.io)
[![Database](https://img.shields.io/badge/Database-SQLite-003B57?style=for-the-badge&logo=sqlite)](https://sqlite.org)
[![Documentation](https://img.shields.io/badge/API_Docs-Swagger-85EA2D?style=for-the-badge&logo=swagger)](http://localhost:8080/swagger/index.html)

This is a high-performance, production-ready backend service designed for a library CRUD application and a specialized URL Processor service. Built with **Golang** using the **Layered Architecture (Service-Repository Pattern)**.

---

## 🏛️ Layered Architecture
This project strictly implements a 3-tier Layered Architecture to decouple concerns and optimize testability:
1. **Presentation Layer (Handlers):** Manages HTTP routing, request binding, and JSON input validations.
2. **Business Logic Layer (Services):** Core domain orchestration and validation rules.
3. **Data Access Layer (Repositories):** Complete database abstraction using GORM.

```text
  [ Client Request ]
         │
         ▼
 ┌───────────────┐
 │   Handlers    │ ◄─── (CORS & Logger Middleware)
 └───────┬───────┘
         │
         ▼
 ┌───────────────┐
 │   Services    │
 └───────┬───────┘
         │
         ▼
 ┌───────────────┐
 │ Repositories  │
 └───────┬───────┘
         │
         ▼
 ┌───────────────┐
 │  SQLite DB    │
 └───────────────┘
```

### 📁 Project Structure
```text
backend/
├── config/             # Database initialization and connection
│   └── database.go
├── docs/               # Auto-generated Swagger specifications
├── handlers/           # HTTP presentation layer
│   ├── book_handler.go
│   └── url_handler.go
├── models/             # Domain and validation schemas
│   ├── book.go
│   └── url.go
├── repositories/       # Data Access Object pattern
│   └── book_repository.go
├── services/           # Business logic orchestration
│   ├── book_service.go
│   ├── url_service.go
│   └── url_service_test.go
├── go.mod
├── go.sum
└── main.go
```

---

## 🚀 Quick Start

### ⚙️ Prerequisites
- Go 1.20 or higher installed.

### 🔌 Running the Server
1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Download project dependencies:
   ```bash
   go mod download
   ```
3. Run the application:
   ```bash
   go run main.go
   ```
4. The server will start on port `8080` with SQLite auto-migrations running instantly (creates `library.db`).

---

## 🧪 Running Unit Tests
A robust set of table-driven tests is provided for the URL processing engine:
```bash
go test -v ./services/...
```

---

## 📝 API Endpoints & Usage

### 📖 1. Library Book Management

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **GET** | `/api/books` | Retrieve all books |
| **POST** | `/api/books` | Add a new book (with model bindings) |
| **GET** | `/api/books/:id` | Get book details by ID |
| **PUT** | `/api/books/:id` | Update existing book by ID |
| **DELETE** | `/api/books/:id` | Delete book entry by ID |

#### **Add Book Schema (POST /api/books):**
```json
{
  "title": "The Clean Architecture",
  "author": "Robert C. Martin",
  "year": 2017
}
```

---

### 🔗 2. URL Cleanup & Redirection Service

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **POST** | `/api/url-process` | Process a URL (canonical, redirection, or all) |

#### **Example Requests & Expected Responses:**

* **All (Canonical & Redirection):**
  - **Request Body:**
    ```json
    {
      "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
      "operation": "all"
    }
    ```
  - **Response Body:**
    ```json
    {
      "processed_url": "https://www.byfood.com/food-experiences"
    }
    ```

* **Canonical Only:**
  - **Request Body:**
    ```json
    {
      "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
      "operation": "canonical"
    }
    ```
  - **Response Body:**
    ```json
    {
      "processed_url": "https://BYFOOD.com/food-EXPeriences"
    }
    ```

---

## 🌐 Interactive API Reference (Swagger)
An interactive documentation console is automatically mounted. Run the server and navigate to:
👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**
