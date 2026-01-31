# Kasir Go API

A simple REST API for POS (Point of Sale) system with Products and Categories management.

## Quick Start

### Prerequisites
- Go 1.25.5+
- PostgreSQL database

### Setup
1. Create `.env` file with your database connection:
```env
PORT=8080
DB_CONN=postgresql://username:password@host:port/database?sslmode=require
```

### Run
```bash
go run main.go
```

### Verify
```bash
curl http://localhost:8080/health
```

## API Endpoints

### Products

| Method | URL | Description |
|--------|-----|-------------|
| GET | `/api/products` | List all products |
| POST | `/api/products` | Create new product |
| GET | `/api/products/{id}` | Get product by ID |
| PUT | `/api/products/{id}` | Update product |
| DELETE | `/api/products/{id}` | Delete product |

### Categories

| Method | URL | Description |
|--------|-----|-------------|
| GET | `/api/categories` | List all categories |
| POST | `/api/categories` | Create new category |
| GET | `/api/categories/{id}` | Get category by ID |
| PUT | `/api/categories/{id}` | Update category |
| DELETE | `/api/categories/{id}` | Delete category |

### System

| Method | URL | Description |
|--------|-----|-------------|
| GET | `/health` | Health check |

## Data Models

### Product
```json
{
  "id": 1,
  "name": "Es Kelapa",
  "price": 10000,
  "stock": 100
}
```

### Category
```json
{
  "id": 1,
  "name": "Snack",
  "description": "Various snack items"
}
```

## Usage Examples

### Create Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Es Kelapa", "price": 10000, "stock": 100}'
```

### Get All Products
```bash
curl http://localhost:8080/api/products
```

### Update Product
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Es Kelapa Premium", "price": 12000, "stock": 80}'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8080/api/products/1
```

### Create Category
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Snack", "description": "Various snack items"}'
```

### Get All Categories
```bash
curl http://localhost:8080/api/categories
```

## Error Handling

| Status Code | Description |
|-------------|-------------|
| 400 | Invalid request body or ID |
| 404 | Resource not found |
| 405 | Method not allowed |
| 500 | Internal server error |

Error response format:
```json
{"error": "produk tidak ditemukan"}
```

## Architecture

- **Handlers**: HTTP request/response handling
- **Services**: Business logic layer
- **Repositories**: Database operations
- **Models**: Data structures

## Technology Stack

- **Go 1.25.5**
- **PostgreSQL**
- **pgx v5** (PostgreSQL driver)
- **golang-migrate** (Database migrations)

## Deployment Notes

### Railway Deployment
- Use HTTPS URLs for API calls
- Environment: `DB_CONN=postgresql://...&prefer_simple_protocol=true`

### Bruno Testing
API testing collections are available in `bruno/` directory for both local and Railway environments.

## Database Schema

- **Categories**: `id`, `name`, `description`, timestamps
- **Products**: `id`, `name`, `price`, `stock`, `category_id`, timestamps
- Foreign key: `products.category_id` references `categories.id`