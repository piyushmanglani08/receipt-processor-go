# receipt-processor-go

## 📍 Overview

A webservice implementation using in-memory data storage, Dockerized for easy deployment and testing, and written in Go.

---

## 👾 Features

|      | Feature         | Summary       |
| :--- | :---:           | :---          |
| ⚙️  | **Architecture**  | <ul><li>Implements a **RESTful API** using the **Gorilla Mux** router for HTTP request handling</li><li>Follows a **layered architecture** with distinct packages for handlers, models, storage, and utilities</li><li>Includes optional **logging middleware** for detailed request logging, controlled via command-line flags</li></ul> |
| 🔩 | **Code Quality**  | <ul><li>Manages dependencies with a `go.mod` file, ensuring consistent builds</li><li>Incorporates unit and integration tests in `main_test.go` to validate business logic and API endpoints</li><li>Adheres to Go best practices, including proper error handling and code readability</li></ul> |

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
    │     └── points.go
    ├── Dockerfile
    ├── .gitignore
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── main_test.go
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




### 🤖 Usage
Run receipt-processor-gog using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run main.go
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker run -it receipt-processor-go
```


### 🧪 Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```
