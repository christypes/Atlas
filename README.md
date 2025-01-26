# Atlas - Distributed URL Shortener

**Atlas** is a highly scalable, fault-tolerant URL shortener with real-time analytics, global low-latency access, and advanced security features. Built to demonstrate distributed systems expertise.

## Features
- **URL Shortening**: Generate short URLs with custom aliases.
- **Real-Time Analytics**: Track clicks, geolocation, and devices.
- **Scalability**: Sharded database, distributed caching, and load balancing.
- **Fault Tolerance**: Database replication and automatic failover.
- **Security**: Rate limiting, JWT authentication, and DDoS protection.

## Tech Stack
### Frontend
- **Framework**: React
- **State Management**: Redux
- **Styling**: Tailwind CSS
- **Deployment**: Vercel

### Backend
- **Language**: Go (Gin/Fiber) or Java (Spring Boot)
- **Database**: PostgreSQL, Cassandra/ScyllaDB
- **Caching**: Redis
- **Messaging**: Kafka
- **API**: REST + gRPC
- **Deployment**: Docker + Kubernetes
- **Monitoring**: Prometheus + Grafana

## Getting Started
1. Clone the repo:
   ```bash
   git clone https://github.com/your-username/atlas.git
   ```
2. Set up the backend:
    ```cd backend
    go run main.go
    ```
3. Setup the frontend:
    ```cd frontend
    npm install
    npm start
    ```
4. Deploy
    ```
    ./scripts/deploy.sh
    ```
