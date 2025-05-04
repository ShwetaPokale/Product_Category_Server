# Product Category Server

A Go-based server application for managing product categories with PostgreSQL database integration.

## Project Structure

```
Product_Category_Server/
├── src/
│   ├── config/
│   │   └── database.go         # Database configuration
│   ├── controllers/
│   │   └── login/
│   │       └── login.controllers.go  # Login controller
│   ├── db/
│   │   └── migrations/
│   │       └── 001_create_users_table.sql  # Database migrations
│   ├── middlewares/
│   │   └── auth/
│   │       └── auth.middleware.go  # Authentication middleware
│   ├── models/
│   │   └── user/
│   │       ├── user.model.go   # User model
│   │       └── user.repository.go  # User repository
│   ├── routes/
│   │   └── login/
│   │       └── login.routes.go  # Login routes
│   └── main.go                 # Application entry point
├── go.mod                      # Go module file
└── README.md                   # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/lenovo/Product_Category_Server.git
   cd Product_Category_Server
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up PostgreSQL:
   - Create a database named `product_catalog` (or set your preferred name in environment variables)
   - Run the migration:
     ```bash
     psql -U postgres -d product_catalog -f src/db/migrations/001_create_users_table.sql
     ```

4. Configure environment variables (or use defaults):
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=your_password
   export DB_NAME=product_catalog
   ```

## Running the Application

1. Start the server:
   ```bash
   go run src/main.go
   ```
   The server will start on port 8080.

## API Endpoints

### Authentication

- `POST /api/login`
  - Request body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "message": "Login successful",
      "token": "jwt-token"
    }
    ```

- `POST /api/logout`
  - Response:
    ```json
    {
      "message": "Logout successful"
    }
    ```

## Development

### Project Structure Details

- `config/`: Configuration management
  - `database.go`: Database connection configuration

- `controllers/`: Request handlers
  - `login/`: Authentication controllers
    - `login.controllers.go`: Login and logout handlers

- `db/`: Database related files
  - `migrations/`: SQL migration files
    - `001_create_users_table.sql`: Users table schema

- `middlewares/`: HTTP middleware
  - `auth/`: Authentication middleware
    - `auth.middleware.go`: JWT authentication middleware

- `models/`: Data models and repositories
  - `user/`: User-related models
    - `user.model.go`: User data structure
    - `user.repository.go`: Database operations for users

- `routes/`: Route definitions
  - `login/`: Authentication routes
    - `login.routes.go`: Login and logout route handlers

### Database Schema

#### Users Table
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 