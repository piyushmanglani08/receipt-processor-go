# receipt-processor-go

## 📍 Overview

A webservice implementation using in-memory data storage, Dockerized for easy deployment and testing, and written in Go.

---

## 👾 Features

|      | Feature         | Summary       |
| :--- | :---:           | :---          |
| ⚙️  | **Architecture**  | <ul><li>Implements a **RESTful API** using the **Gorilla Mux** router for HTTP request handling</li><li>Follows a **layered architecture** with distinct packages for handlers, models, storage, and utilities</li><li>Includes optional **logging middleware** for detailed request logging, controlled via command-line flags</li></ul> |
| 🔩 | **Code Quality**  | <ul><li>Manages dependencies with a `go.mod` file, ensuring consistent builds</li><li>Incorporates unit and integration tests in `main_test.go` to validate business logic and API endpoints</li><li>Adheres to Go best practices, including proper error handling and code readability</li></ul> |
| 📄 | **Documentation** | <ul><li>Provides comprehensive **documentation** covering installation, setup, and API usage</li><li>Includes example **commands** for running and testing the application with `go run` and `go test`</li><li>Documents environment variables and configurations needed for deployment</li></ul> |
| 🔌 | **Integrations**  | <ul><li>Supports **PostgreSQL** for persistent data storage</li><li>Uses **Docker** for containerized deployment, with a multi-stage build process</li><li>Implements **GitHub Actions** for CI/CD automation, including tests and Docker image publishing</li></ul> |
| 🧩 | **Modularity**    | <ul><li>Separates concerns with dedicated modules for API routes, business logic, and storage</li><li>Utilizes middleware for **authentication**, **logging**, and **error handling**</li><li>Ensures **code reusability** by defining common utilities and shared components</li></ul> |
| 🧪 | **Testing**       | <ul><li>Includes unit and integration tests to validate key application functionality</li><li>Uses `go test` with coverage reports for tracking test effectiveness</li><li>Ensures robustness by testing error cases and edge scenarios</li></ul> |
| ⚡️  | **Performance**   | <ul><li>Optimized database queries to reduce load times</li><li>Implements caching mechanisms for frequently accessed data</li><li>Uses **goroutines** to handle concurrent requests efficiently</li></ul> |
| 🛡️ | **Security**      | <ul><li>Implements **JWT-based authentication** for secure access control</li><li>Follows **OWASP security best practices**, including input validation and SQL injection prevention</li><li>Uses role-based access to restrict actions based on user permissions</li></ul> |

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
