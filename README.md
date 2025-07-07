# Go Microservices Demo

A simple event-driven microservices demo built with Go, featuring an e-commerce backend with Kafka for inter-service communication.

## Architecture

The system consists of three microservices:

- **order-service**: REST API that receives orders and publishes events to Kafka
- **inventory-service**: Consumes order events and updates inventory
- **notification-service**: Consumes order events and logs notifications

All services communicate asynchronously via Kafka events.

## Prerequisites

- Go 1.22+
- Docker and Docker Compose
- Git

## Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Seiya-P/go-microservices-demo.git
   cd go-microservices-demo
   ```

2. **Start all services:**
   ```sh
   docker-compose up --build
   ```

3. **Test the system:**
   ```sh
   curl -X POST http://localhost:8080/orders \
     -H "Content-Type: application/json" \
     -d '{"id":"order-123","item":"widget","amount":5}'
   ```

### Local Development

1. **Start Kafka and Zookeeper:**
   ```sh
   docker-compose up zookeeper kafka
   ```

2. **Run services locally:**
   ```sh
   # Terminal 1
   cd order-service && go run main.go
   
   # Terminal 2
   cd inventory-service && go run main.go
   
   # Terminal 3
   cd notification-service && go run main.go
   ```

## API Documentation

### Order Service

**Endpoint:** `POST /orders`

**Request Body:**
```json
{
  "id": "order-123",
  "item": "widget",
  "amount": 5
}
```

**Response:** `201 Created` on success

**Example:**
```sh
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"id":"order-123","item":"widget","amount":5}'
```

## Project Structure

```
go-microservices-demo/
├── order-service/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── inventory-service/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── notification-service/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── docker-compose.yml
├── .github/workflows/ci.yml
└── README.md
```

## Services

### Order Service
- **Port:** 8080
- **Purpose:** Receives orders via REST API
- **Events:** Publishes `OrderCreated` events to Kafka topic `orders`

### Inventory Service
- **Purpose:** Consumes order events and updates inventory
- **Events:** Listens to Kafka topic `orders`
- **Storage:** In-memory inventory (widget: 100, gadget: 100)

### Notification Service
- **Purpose:** Consumes order events and logs notifications
- **Events:** Listens to Kafka topic `orders`

## Development

### Building Services
```sh
# Build all services
cd order-service && go build
cd ../inventory-service && go build
cd ../notification-service && go build
```

### Running Tests
```sh
# Test each service
cd order-service && go test
cd ../inventory-service && go test
cd ../notification-service && go test
```

### Docker Builds
```sh
# Build Docker images
docker build -t order-service ./order-service
docker build -t inventory-service ./inventory-service
docker build -t notification-service ./notification-service
```

## CI/CD

The project includes GitHub Actions workflows that:
- Build and test all Go services
- Run code linting and formatting checks
- Build Docker images for each service

## Kafka Topics

- **orders**: Contains order events published by order-service

## Monitoring

Check service logs:
```sh
# Docker Compose logs
docker-compose logs order-service
docker-compose logs inventory-service
docker-compose logs notification-service

# Individual service logs
docker-compose logs -f order-service
```

## Troubleshooting

### Kafka Connection Issues
- Ensure Kafka and Zookeeper are running: `docker-compose up zookeeper kafka`
- Check Kafka is accessible at `localhost:9092`

### Service Communication
- Verify all services are running: `docker-compose ps`
- Check service logs for connection errors

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE). 