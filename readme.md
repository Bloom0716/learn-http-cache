# Learn HTTP Cache

This project demonstrates the use of HTTP caching mechanisms with a Go backend and Nginx as a reverse proxy.
It is designed as a learning resource for understanding HTTP caching concepts and is intended to be used as part of a new graduate training program.

## Features

- **HTTP Caching**: Demonstrates `ETag`, `Cache-Control`, and Nginx caching.
- **API Endpoints**:
  - `/api/health`: Health check endpoint.
  - `/api/now`: Returns the current time.
  - `/api/now/no-store`: Returns the current time with `Cache-Control: no-store`.
  - `/api/users/:id`: Fetches user details with `ETag` support.
  - `/api/users/create`: Creates a new user.

## Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/Bloom0716/learn-http-cache.git
   cd learn-http-cache
   ```

2. Build and start the services:
   ```sh
   make up
   ```

3. Access the application:
   - Backend: `http://localhost:8080`
   - Nginx Proxy: `http://localhost:8081`

## Scripts

- `script/create_user.sh`: Creates a new user.
- `script/get_user.sh`: Fetches user details with `ETag`.
- `script/health.sh`: Checks the health of the application.
- `script/now.sh`: Fetches the current time.
- `script/now_no_store.sh`: Fetches the current time with `Cache-Control: no-store`.

## Technologies Used

- **Go**: Backend implementation.
- **Nginx**: Reverse proxy and caching.
- **PostgreSQL**: Database for user data.
- **Docker**: Containerization.
- **Air**: Live reload for Go development.
