# receipt-processor-go

## 📍 Overview

A webservice implementation using in-memory data storage, Dockerized for easy deployment and testing, and written in Go.

---

## 📁 Project Structure

```sh
└── receipt-processor-go/
    ├── handlers
    │     ├── getPoints.go
    │     └── process.go
    ├── models
    │     └── receipt.go
    ├── storage
    │     └── store.go
    ├── utils
    │     ├── response.go
    │     └── points.go
    ├── Dockerfile
    ├── .gitignore
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── main.go
```

## 🚀 Getting Started

### ☑️ Prerequisites

Before getting started with docker-gs-ping, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules
- **Container Runtime:** Docker


### ⚙️ Installation

Install docker-gs-ping using one of the following methods:

**Build from source:**

1. Clone the docker-gs-ping repository:
```sh
❯ git clone https://github.com/piyushmanglani08/receipt-processor-go.git
```

2. Navigate to the project directory:
```sh
❯ cd receipt-processor-go
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go mod tidy
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker build -t receipt-processor-go .


```

### 🤖 Running the application
Run receipt-processor-gog using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run main.go
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker run -p 8080:8080 -it receipt-processor-go
```


### 🧪 Testing API Usage

You can test the API using **cURL** or **Postman**.

### 1️⃣ Process a Receipt
#### **Endpoint:** `POST /receipts/process`
#### **Request:**
```sh
curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d '{
          "retailer": "Target",
          "purchaseDate": "2022-01-01",
          "purchaseTime": "13:01",
          "total": "35.35",
          "items": [
              { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
              { "shortDescription": "Emils Cheese Pizza", "price": "12.25" },
              { "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },
              { "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },
              { "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00" }
          ]
      }'
```
#### **Response:**
```json
{
    "id": "674d7acc-9eb7-52b9-85c3-4384c200bcbe"
}
```

### 2️⃣ Retrieve Points for a Receipt
#### **Endpoint:** `GET /receipts/{id}/points`
#### **Request:**
```sh
curl -X GET "http://localhost:8080/receipts/674d7acc-9eb7-52b9-85c3-4384c200bcbe/points"
```
#### **Response:**
```json
{
    "points": 28
}
```
