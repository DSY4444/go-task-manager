# Go Task Manager

This is a simple task manager REST API built with Go, Gorilla Mux, GORM, PostgreSQL, and JWT authentication.

## Setup

1. Install Go and PostgreSQL
2. Clone the repo and install dependencies:
   ```sh
   go mod tidy
   ```
3. Run PostgreSQL using Docker:
   ```sh
   docker run --name taskdb -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=taskdb -p 5432:5432 -d postgres
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

- `POST /login` - Get JWT Token
- `GET /tasks` - Get all tasks (Auth required)
- `POST /tasks` - Create a new task (Auth required)
- `PUT /tasks/{id}` - Update a task (Auth required)
- `DELETE /tasks/{id}` - Delete a task (Auth required)
   