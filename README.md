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

atlas/
├── frontend/                  # Frontend code
│   ├── public/                # Static assets
│   ├── src/                   # React source code
│   │   ├── components/        # Reusable components
│   │   ├── pages/             # Page components
│   │   ├── store/             # Redux state management
│   │   ├── utils/             # Utility functions
│   │   ├── App.js             # Main app component
│   │   └── index.js           # Entry point
│   ├── tailwind.config.js     # Tailwind CSS config
│   └── package.json           # Frontend dependencies
│
├── backend/                   # Backend code
│   ├── api/                   # API handlers
│   ├── config/                # Configuration files
│   ├── db/                    # Database models and migrations
│   ├── middleware/            # Custom middleware (e.g., auth, rate limiting)
│   ├── services/              # Business logic
│   ├── utils/                 # Utility functions
│   ├── main.go                # Entry point (Go) / Application.java (Java)
│   └── go.mod                 # Go dependencies / pom.xml (Java)
│
├── scripts/                   # Deployment and utility scripts
│   ├── deploy.sh              # Deployment script
│   └── migrate.sh             # Database migration script
│
├── docker/                    # Dockerfiles
│   ├── frontend.Dockerfile    # Frontend Dockerfile
│   ├── backend.Dockerfile     # Backend Dockerfile
│   └── docker-compose.yml     # Docker Compose for local dev
│
├── docs/                      # Documentation
│   ├── architecture.md        # High-level architecture
│   └── api.md                 # API documentation
│
├── .github/                   # GitHub workflows
│   └── workflows/             # CI/CD pipelines
│       ├── frontend-ci.yml    # Frontend CI
│       └── backend-ci.yml     # Backend CI
│
├── .gitignore                 # Git ignore rules
├── README.md                  # Project overview
└── LICENSE                    # Project license