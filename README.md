# Application API

A simple Go HTTP API that returns back whatever valid JSON send to it. 

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
./run.sh
```

This will start 3 instances on ports 8080, 8081, and 8082 for load balancing testing.

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
  }
}
```

### GET /health

Check if the service is alive and well:

```bash
curl http://localhost:8080/health
```

```json
{
  "status": "healthy"
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
├── run.sh               # Script to run multiple instances
├── config/
│   └── config.go        # Handles environment variables
├── handlers/
│   └── handlers.go      # The actual API endpoints
├── middleware/
│   └── middleware.go    # Logging and other middleware
└── README.md            # This file
```

## Integration with Routing API

This application API is designed to work with the routing API for load balancing:

1. **Start multiple instances:**
   ```bash
   ./run.sh
   ```
   This starts instances on ports 8080, 8081, and 8082.

2. **Start the routing API** (in the routing-api directory):
   ```bash
   ./run.sh
   ```

3. **Test load balancing:**
   ```bash
   # Requests to routing API (port 3000) will be distributed across application APIs
   curl -X POST http://localhost:3000/testapi \
     -H "Content-Type: application/json" \
     -d '{"message": "test load balancing"}'
   ```