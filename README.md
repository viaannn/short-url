# short-url
[![Go Version](https://img.shields.io/badge/Go-1.21-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

## 📌 Overview
A brief description of the project, what it does, and its key features.

## 🚀 Features
- ✅ API to Create new short url
- ✅ API to Redirect to full url

## 📦 Directory Structure

```
.
├── cmd/            # Entry points for the application (webserver)
├── internal/       # Private application logic (service, repository layer)
├── config/         # Configuration files (generic variable)
├── migrations/     # Database migrations
├── tests/          # Unit tests
├── Dockerfile      # Docker configuration
├── Makefile        # Build automation
└── README.md       # Project documentation
```

## ⚡ Installation

### **1. Clone the repository**
```sh
https://github.com/viaannn/short-url-api.git
cd short-url-api
```

### **2. Install dependencies**
```sh
go mod tidy
```

### **3. Set up environment variables**
Copy the example environment file and update it as needed:
```sh
cp .env.example .env
```

### **4. Run the application**
```sh
go run cmd/webserver/main.go
```

Or using `Makefile`:
```sh
make run
```

## 🛠️ Configuration

Environment variables required:

```ini
PORT=8080
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=secret
DB_NAME=mydb
```

## 🧪 Running Tests

### **Run all tests**
```sh
go test ./test
```

### **Run tests with coverage**
```sh
go test -cover ./test
``` 

## 📦 Docker Support

### **Build Docker Image**
```sh
docker build -t short-url-api .
```

### **Run Docker Container**
```sh
docker run -p 8080:8080 --env-file .env short-url-api
```

## 🚀 Deployment

### **Using Docker Compose**
```sh
docker-compose up -d
```

### **Using GitHub Actions**
See `.github/workflows/deploy.yml` for CI/CD automation.

## 🐜 License

This project is licensed under the [MIT License](LICENSE).

---

🌟 **Don't forget to give this project a star if you find it useful!** 🌟