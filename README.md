# Product Aggregator API

A Golang-based API that aggregates product data from multiple providers, normalizes responses, and exposes a unified search endpoint.
## 🚀 Features

- Fetch data from multiple providers concurrently
- Normalize different provider response formats
- Deduplicate products by SKU
- Sort products by price
- Timeout handling (800ms per provider)
- Track failed providers
- Store search history
- Health check endpoint

---

## 🛠️ Tech Stack

- Go (Golang)
- Gin (HTTP framework)

---

## ⚙️ Setup Instructions

### 1. Clone repository

git clone <your-repo-url>
cd product-api

### 2. Install dependencies

go mod tidy

### 3. Run the application

go run main.go

---

## 🌐 API Endpoints

---

### 🔍 Search Products

GET /search?q=<query>

Example:

http://localhost:8082/search?q=mouse

Response:

```json
{
  "query": "mouse",
  "products": [...],
  "failed_providers": ["provider_c"]
}
