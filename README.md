# Go Backend Project

A small Go HTTP API built with [Gorilla Mux](https://github.com/gorilla/mux) and PostgreSQL. It is intended as a simple starting point for experimenting with Go backend development.

## Requirements

- Go 1.23.2 or newer
- Docker and Docker Compose (for the local PostgreSQL instance)

## Setup

1. Clone the repository and enter the project directory:

   ```bash
   git clone https://github.com/gunash-portfolio/go_backend_project.git
   cd go_backend_project
   ```

2. Download Go dependencies:

   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root. These values are used by both the application and Docker Compose:

   ```dotenv
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=go_backend
   ```

   Keep local secrets in `.env` and do not commit that file.

4. Start PostgreSQL:

   ```bash
   docker compose up -d postgres
   ```

## Run the API

Start the server with:

```bash
go run main.go
```

The API listens on `http://localhost:8080`. The database must be running and the `.env` values must be set before starting the application.

## Endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/` | Returns a welcome message |
| `GET` | `/api/hello` | Returns a hello message |

Example requests:

```bash
curl http://localhost:8080/
curl http://localhost:8080/api/hello
```

## Stop the Local Database

```bash
docker compose down
```

To remove the database volume and its data as well:

```bash
docker compose down -v
```

## Tech Stack

- Go 1.23.2
- Gorilla Mux v1.8.1
- PostgreSQL 15
- `github.com/lib/pq`

