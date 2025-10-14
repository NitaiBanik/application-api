# Application API

A simple Go HTTP API that echoes back whatever JSON you send to it. Pretty handy for testing things out or when you need a basic API that just returns what you put in.


## Getting started

You'll need Go 1.21 or newer installed on your machine.

1. **Clone this repo:**
   ```bash
   git clone <your-repo-url>
   cd application-api
   ```

2. **Get the dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run it:**
   ```bash
   go run main.go
   ```

4. **Test api:**
   ```bash
   curl -X POST http://localhost:8080/testapi \
     -H "Content-Type: application/json" \
     -d '{"message": "Hello World"}'
   ```

### Quick start with testing
```bash
./run-and-test.sh
```

### Build it
```bash
go build -o application-api main.go
```

### Run tests
```bash
go test ./...
```

Run specific tests:
```bash
go test ./handlers   
go test -run TestAPI
```

## How the service works

### POST /testapi

Send it any valid JSON and it'll send the same thing back to you.

```bash
curl -X POST http://localhost:8080/testapi \
  -H "Content-Type: application/json" \
  -d '{"key": "value", "number": 42, "array": [1, 2, 3]}'
```

You'll get back:
```json
{
  "data": {
    "key": "value",
    "number": 42,
    "array": [1, 2, 3]
  },
  "timestamp": "2024-01-01T12:00:00Z"
}
```

### GET /health

Check if the service is alive and well:

```bash
curl http://localhost:8080/health
```

```json
{
  "status": "healthy",
  "service": "application-api",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Configuration
Copy the example env file and modify it:
```bash
cp env.example .env
# edit .env with your values
```

Or set them directly:
```bash
PORT=3000 go run main.go
```

## Project structure

```
application-api/
├── main.go              # Where it all starts
├── go.mod               # Dependencies
├── env.example          # Environment variables template
├── run-and-test.sh      # Script to run and test the API
├── config/
│   └── config.go        # Handles environment variables
├── handlers/
│   └── handlers.go      # The actual API endpoints
├── middleware/
│   └── middleware.go    # Logging and other middleware
└── README.md            # This file
```