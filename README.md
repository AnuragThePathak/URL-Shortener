# Minly — URL Shortener

A small full-stack project built in 2022 to practice end-to-end application ownership: backend API design, interface-driven architecture, database integration, and automated deployment.

**Live:** [mly.vercel.app](https://mly.vercel.app)

---

## Overview

Minly is a URL shortener. You give it a long URL, it returns a short one. Visiting the short URL redirects to the original.

### Why build this?

At the time, I wanted to:

- Write a backend in Go without reaching for a heavy framework
- Practice interface-driven design with clear component boundaries
- Work through the full cycle: design → build → deploy → maintain
- Set up CI/CD pipelines that deploy on release, not on every merge

This project is intentionally small in scope. It's not trying to demonstrate distributed systems, caching layers, or high-availability patterns. It's a clean, working service that covers the fundamentals with good structure.

---

## Architecture

The backend follows a layered architecture with clear separation between HTTP handling, business logic, and data persistence.

```
┌─────────────────────────────────────────────────────────────┐
│                      HTTP Request                           │
└─────────────────────────┬───────────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────────┐
│  endpoints/url_endpoints.go                                 │
│  - Parses requests, returns responses                       │
│  - Depends on UrlService interface                          │
└─────────────────────────┬───────────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────────┐
│  api/url.go (UrlService)                                    │
│  - URL validation (parse + DNS lookup)                      │
│  - Duplicate detection                                      │
│  - ID generation via Sonyflake                              │
│  - Depends on UrlStore interface                            │
└─────────────────────────┬───────────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────────┐
│  api/mongodb/url_store.go                                   │
│  - Concrete storage implementation                          │
│  - Translates service types to MongoDB documents            │
└─────────────────────────┬───────────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                        MongoDB                              │
└─────────────────────────────────────────────────────────────┘
```

### Key Design Decisions

- **No framework on the backend.** The HTTP server uses Go's `net/http` with chi for routing. Middleware and request handling are explicit.

- **Interface-driven design.** The `UrlStore` interface abstracts storage. The MongoDB implementation sits behind it. Swapping storage layers doesn't touch business logic.

- **ID generation with Indigo.** Short codes are generated using `osamingo/indigo`, a Sonyflake-based generator. This produces unique IDs without database coordination.

- **Validation before persistence.** URLs are parsed and DNS-checked before storing. Invalid or unreachable URLs fail fast.

- **Dependency injection.** Components receive dependencies at construction time. The `main` function acts as the composition root.

- **Graceful shutdown.** The server listens for OS signals and shuts down cleanly, giving in-flight requests time to complete.

- **Structured logging.** `zap` provides environment-aware logging (development vs production formats).

---

## Project Structure

```
backend/
├── main.go              # Composition root, wires all dependencies
├── config.go            # Environment and server configuration
├── api/
│   ├── url.go           # UrlService interface and implementation
│   ├── url_test.go      # Service tests with mocked store
│   └── mongodb/
│       └── url_store.go # UrlStore implementation
├── endpoints/
│   └── url_endpoints.go # HTTP handlers
├── server/
│   ├── server.go        # HTTP server with graceful shutdown
│   └── endpoints.go     # Request/response utilities
├── common/
│   └── constants.go     # Shared constants
└── Dockerfile           # Multi-stage build

nextjs/                  # Minimal frontend (not the focus)
```

---

## Tech Stack

| Layer    | Technology              |
|----------|-------------------------|
| Backend  | Go (standard library + chi) |
| Database | MongoDB                 |
| Frontend | Next.js                 |
| CI/CD    | GitHub Actions          |
| Deploy   | Heroku (backend), Vercel (frontend) |

---

## CI/CD Philosophy

The CI/CD setup reflects a clear philosophy: **merge to master should not automatically deploy.** Deployments happen only when a GitHub Release is explicitly published.

### Continuous Integration (every push/PR to master)

- **Go workflow:** Build and test the backend
- **Node.js workflow:** Build frontend across Node 14.x and 16.x
- **CodeQL:** Security scanning for Go and JavaScript

### Continuous Deployment (on GitHub Release only)

- **Backend:** Docker image built and pushed to Heroku Container Registry
- **Frontend:** Deployed to Vercel production

This separation provides:
- Clear audit trail of deployments
- Prevention of accidental deploys from incomplete work
- Ability to batch changes into a single release
- Easy rollback targets

---

## Running Locally

```bash
cd backend

export DB_URL="mongodb://..."
export DB_NAME="minly"

go run .
```

The server starts on port `8080` by default. Optional environment variables:

| Variable      | Description                | Default |
|---------------|----------------------------|---------|
| `PORT`        | Server port                | 8080    |
| `ENV`         | `production` or `development` | development |
| `TLS_ENABLED` | Enable HTTPS               | false   |
| `TLS_CERT_PATH` | Path to TLS certificate  | —       |
| `TLS_KEY_PATH`  | Path to TLS key          | —       |

---

## Testing

```bash
cd backend
go test -v ./...
```

Tests use mock implementations of the `UrlStore` interface, allowing the service layer to be tested without a live database.

---

## Scope & Limitations

This project is a learning exercise, not a production system. It does not attempt to solve:

- Rate limiting or abuse prevention
- Analytics or click tracking
- Custom short codes or user accounts
- Horizontal scaling or caching
- URL expiration or cleanup

These are valid concerns for a real URL shortener, but they weren't the goal here.

---

## What This Project Demonstrates

Despite its small scope, this project demonstrates several engineering practices:

- **Interface-based design** for testability and decoupling
- **Graceful shutdown** with OS signal handling
- **Clear CI/CD separation** between integration and deployment
- **Centralized request handling** for consistent error responses
- **Configuration through environment variables** (twelve-factor)
- **Multi-stage Docker builds** for minimal production images
- **Unit tests with mocks** for isolated testing

---

## License

[MIT](LICENSE)
