# Go Simple CRUD API 3

A simple RESTful API built with Go, PostgreSQL, and standard libraries. This project implements a layered architecture (Controller/Handler, Service, Repository, Model) for managing Categories, Products, Transactions (Checkout), and Sales Reports.

## Live Demo

https://go-simple-crud-3-production.up.railway.app/

## Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Configuration:** Viper
- **Driver:** lib/pq

## Project Structure

```
.
├── database/           # Database connection & initialization
├── handlers/           # HTTP Handlers (Controllers)
├── models/             # Data structures (Entities)
├── repositories/       # Database access layer (CRUD logic)
├── services/           # Business logic layer
├── .env                # Environment variables (Configuration)
├── go.mod              # Go dependencies
└── main.go             # Application entry point & Routing
```

## API Endpoints

### Categories

| Method | Endpoint             | Description              |
| ------ | -------------------- | ------------------------ |
| GET    | `/api/category`      | Get all categories       |
| GET    | `/api/category/{id}` | Get category by ID       |
| POST   | `/api/category`      | Create new category      |
| PUT    | `/api/category/{id}` | Update existing category |
| DELETE | `/api/category/{id}` | Delete category          |

### Products

| Method | Endpoint            | Description                               |
| ------ | ------------------- | ----------------------------------------- |
| GET    | `/api/product`      | Get all products (includes category name) |
| GET    | `/api/product/{id}` | Get product by ID (includes category name)|
| POST   | `/api/product`      | Create new product                        |
| PUT    | `/api/product/{id}` | Update existing product                   |
| DELETE | `/api/product/{id}` | Delete product                            |

### Transactions

| Method | Endpoint        | Description                         |
| ------ | --------------- | ----------------------------------- |
| POST   | `/api/checkout` | Create new transaction (multi-item) |

### Reports

| Method | Endpoint                                                 | Description                    |
| ------ | -------------------------------------------------------- | ------------------------------ |
| GET    | `/api/report`                                            | Get all-time sales report      |
| GET    | `/api/report?start_date=2026-01-01&end_date=2026-02-01` | Get sales report by date range |
| GET    | `/api/report/hari-ini`                                   | Get today's sales report       |

## API Request & Response Examples

### 1. Create Category

**POST** `/api/category`

```json
{
  "name": "Electronics",
  "description": "Gadgets and devices"
}
```

### 2. Create Product

**POST** `/api/product`

```json
{
  "name": "Smartphone",
  "price": 5000000,
  "stock": 50,
  "category_id": 1
}
```

### 3. Get Product Detail

**GET** `/api/product/1`

Response:

```json
{
  "id": 1,
  "name": "Smartphone",
  "price": 5000000,
  "stock": 50,
  "category_id": 1,
  "category_name": "Electronics"
}
```

### 4. Checkout (Create Transaction)

**POST** `/api/checkout`

```json
{
  "items": [
    { "product_id": 1, "quantity": 2 },
    { "product_id": 3, "quantity": 1 }
  ]
}
```

Response:

```json
{
  "id": 1,
  "total_amount": 15000000,
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Smartphone",
      "quantity": 2,
      "subtotal": 10000000
    },
    {
      "id": 2,
      "transaction_id": 1,
      "product_id": 3,
      "product_name": "Headphones",
      "quantity": 1,
      "subtotal": 5000000
    }
  ]
}
```

### 5. Get Sales Report

**GET** `/api/report/hari-ini`

Response:

```json
{
  "total_revenue": 45000,
  "total_transaksi": 5,
  "produk_terlaris": {
    "nama": "Indomie Goreng",
    "qty_terjual": 12
  }
}
```

## Run Locally

### 1. Prerequisites

- Go 1.22+
- PostgreSQL database

### 2. Setup Environment

Create a `.env` file in the root directory:

```
PORT=8080
DB_CONNECTION=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

## Database Schema

The project uses a relational schema:

- **categories:** `id`, `created_at`, `name`, `description`
- **products:** `id`, `created_at`, `name`, `price`, `stock`, `category_id` (FK)
- **transactions:** `id`, `total_amount`, `created_at`
- **transaction_details:** `id`, `transaction_id` (FK), `product_id` (FK), `quantity`, `subtotal`, `created_at`

> **Note**
> - When fetching products, the application performs a LEFT JOIN with the categories table to provide the `category_name` in the response for better usability.
> - Checkout automatically validates product stock, calculates subtotals, and deducts stock within a database transaction to ensure data consistency.
> - Sales reports aggregate data from transactions and transaction_details with JOIN to products for the best-selling product.
