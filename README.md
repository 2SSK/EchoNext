# EchoNext

A full-stack web application demonstrating end-to-end development skills, from infrastructure deployment to production-ready backend and frontend implementation.

## Tech Stack

- **Backend**: Go with Echo framework, PostgreSQL database
- **Frontend**: Next.js with React and TypeScript
- **Infrastructure**: AWS deployment with CI/CD pipeline
- **DevOps**: Docker containerization, monitoring, and automation

## Features

- RESTful API with health checks and error handling
- Database migrations and connection pooling
- CORS-enabled server with rate limiting
- Structured logging with Zerolog
- Static file serving for SPA routing
- Production-ready Docker builds

## Getting Started

1. Clone the repository
2. Set up environment variables (see `.env.example`)
3. Run database migrations
4. Start the backend: `go run cmd/EchoNext/main.go`
5. Start the frontend: `cd app && pnpm dev`

## Deployment

Deployed on AWS with automated CI/CD pipeline for seamless updates and scaling.