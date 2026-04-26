# Product Aggregator API

A Golang-based backend service that aggregates product data from multiple providers, normalizes responses, and exposes a unified API.

---

## 🚀 Features

- Fetch products from multiple providers concurrently  
- Normalize different provider response formats  
- Deduplicate products by SKU (keeps lowest price)  
- Sort products by price (ascending)  
- Timeout handling (800ms per provider)  
- Graceful handling of provider failures  
- Search history tracking  
- Health check endpoint  

---

## 🛠️ Tech Stack

- Go (Golang)  
- Gin (HTTP framework)  

---

## ⚙️ Setup Instructions

### 1. Prerequisites

Make sure you have the following installed:

- Go (version 1.25 or higher)  

```bash
go version
```

- Git  

---

### 2. Clone the Repository

```bash
git clone https://github.com/hafiskc/product-api.git
cd product-api
```

---

### 3. Install Dependencies

```bash
go mod tidy
```

---

### 4. Run the Application

```bash
go run main.go
```

You should see:

```
Listening and serving HTTP on :8089
```

---

### 5. Verify the Service

```bash
curl http://localhost:8089/health
```

Example response:

```json
{
  "status": "ok",
  "timestamp": "2026-04-26T13:00:00Z"
}
```

---

## 🌐 API Endpoints

### 🔍 Search Products

**GET /search?q=<query>**

Example:

```
http://localhost:8089/search?q=mouse
```

Response:

```json
{
  "data" : {

    "query": "mouse",
    "products": [
      {
        "sku": "A1",
        "name": "Wireless Mouse",
        "price": 19.99,
        "currency": "USD",
        "provider": "provider_a"
      }
    ],
    "failed_providers": ["provider_c"]
  }
}
```

---

### 📜 Get Search History

**GET /history**

Returns the latest 10 search records  

---

### 📄 Get Search by ID

**GET /history/{id}**

Returns details of a specific search  

---

### ❤️ Health Check

**GET /health**

Response:

```json
{
  "status": "ok",
  "timestamp": "..."
}
```

---

## 🧪 Example cURL Commands

```bash
curl "http://localhost:8082/search?q=mouse"
curl "http://localhost:8082/history"
curl "http://localhost:8082/history/<id>"
curl "http://localhost:8082/health"
```

---

## 🧪 Testing

Basic unit tests are implemented for the service layer to verify core functionality such as aggregation, deduplication, and error handling.

### ▶️ Run All Tests

```bash
go test ./...
```

---

## 📂 Project Structure

```
main.go


handler/
service/
provider/
repository/
model/
normalizer/
mockdata/
```

---

## 🧠 Design Decisions

- Used interface-based provider abstraction for flexibility and scalability  
- Implemented concurrency using goroutines and WaitGroup  
- Used context for timeout handling and cancellation  
- Normalization handled in a separate layer for clean architecture  
- In-memory repository used for search history (for simplicity)  

---

## ⚖️ Tradeoffs

- In-memory storage (data lost on restart)  
- No retry mechanism for failed providers  
- No database integration (kept simple for scope)  
- No pagination for results  
- Limited test coverage  

---

## 🚀 Future Improvements

- Add database (PostgreSQL) for persistence  
- Add caching (Redis)  
- Implement retries with exponential backoff  
- Add pagination and filtering   

---

## 📝 Notes

- Default port: **8089**  
- Can be changed in `main.go`:

```go
router.Run(":8089")
```

---

## 👨‍💻 Author

Hafis Muhammad
