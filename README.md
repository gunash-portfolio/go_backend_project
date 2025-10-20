# go_backend_project

A fresh start with Go backend development using the Gorilla Mux router framework.

## About

This is a simple Go web server project using [Gorilla Mux](https://github.com/gorilla/mux) - a powerful HTTP router and URL matcher for building Go web servers.

## Getting Started

### Prerequisites
- Go 1.23.2 or higher

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Available Endpoints

- `GET /` - Home page
- `GET /api/hello` - Hello endpoint

## Tech Stack

- Go 1.23.2
- Gorilla Mux v1.8.1



