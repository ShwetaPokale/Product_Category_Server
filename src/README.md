# Product Catalog API

A RESTful API for managing product catalogs with authentication.

## Project Structure

```
src/
├── middlewares/
│   └── auth/
│       └── auth.middleware.go
├── routes/
│   └── login/
│       └── login.route.go
├── controllers/
│   └── login/
│       └── login.controllers.go
├── test/
│   └── login/
│       └── login.test.go
├── main.go
├── go.mod
└── README.md
```

## Getting Started

1. Install Go 1.21 or later
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Run `go run main.go` to start the server

## API Endpoints

### Authentication

- `POST /api/login` - Login with username and password
- `POST /api/logout` - Logout (requires authentication)

## Testing

Run tests with:
```bash
go test ./...
```

## Dependencies

- github.com/golang-jwt/jwt/v5 - JWT authentication 