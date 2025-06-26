# Book Management REST API

A secure REST API built with Go, PostgreSQL, and JWT authentication for managing books.

## Features

- JWT Authentication (Register/Login)
- CRUD Operations for Books
- PostgreSQL Database
- Swagger API Documentation
- Middleware (Logging, Recovery, Authentication)
- Unit & Integration Tests

## Tech Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL 15+
- **Authentication**: JWT
- **Documentation**: Swagger
- **Testing**: Go's testing package

## Installation

### Prerequisites

1. Install [Go](https://go.dev/dl/)
2. Install [PostgreSQL](https://www.postgresql.org/download/)
3. Install Swag:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
4. The postgres db is "bookdb" in this project
    
## Contributing

- Fork the project
- Create your feature branch (git checkout -b feature/AmazingFeature)
- Commit your changes (git commit -m 'Add some AmazingFeature')
- Push to the branch (git push origin feature/AmazingFeature)
- Open a Pull Request
