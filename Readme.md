# Mobile Number OTP Authentication Service

This is a Golang-based microservice for mobile number/OTP-based authentication. It allows users to register, log in using OTP, resend OTP, and access their details once authenticated.


#### Project structure

msg/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── controllers/
│   └── auth_controller.go
├── middleware/
│   └── middleware.go
├── models/
│   └── user.go
├── services/
│   ├── otp_service.go
│   └── auth_service.go
├── utils/
│   ├── jwt.go
│   └── otp.go
├── routes/
│   └── routes.go
├── .env
├── go.mod
└── go.sum


## Getting Started

### Prerequisites

- Golang version 1.21+
- PostgreSQL and Redis for local development (if not using Docker)


### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/msg.git
cd msg

```

2. Copy .env.example to .env and update the environment variables:

```bash
cp .env.example .env

```

3. The server should now be running on http://localhost:8080.


### Environment Variables

Configure the following environment variables in your .env file:

- DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME: PostgreSQL connection details.
- REDIS_ADDR: Redis server address.
- JWT_SECRET: Secret for JWT generation.


## API Endpoints

### Register User

```bash
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{
  "mobile": "1234567890",
  "name": "John Doe",
  "fingerprint": "device_fingerprint_example"
}'

```

### Request OTP


```bash
curl -X POST http://localhost:8080/login/request-otp -H "Content-Type: application/json" -d '{
  "mobile": "1234567890"
}'

```

### Verify OTP and Get Token

```bash
curl -X POST http://localhost:8080/login/verify-otp -H "Content-Type: application/json" -d '{
  "mobile": "1234567890",
  "otp": "123456"
}'

```

### Resend OTP

```bash
curl -X POST http://localhost:8080/login/resend-otp -H "Content-Type: application/json" -d '{
  "mobile": "1234567890"
}'

```

### Get User Details (Requires Authorization)

```bash
curl -X GET http://localhost:8080/user -H "Authorization: Bearer <JWT_TOKEN>"

```

Replace <JWT_TOKEN> with the token received from the verify-otp endpoint.


## Deployment

For production, modify the .env file with production settings and use a secure way to manage secrets. The Docker Compose configuration includes the following services:

- App Service: Runs the Golang application.
- PostgreSQL Service: Provides persistent storage for user data.
- Redis Service: Caches OTPs and manages rate limits.