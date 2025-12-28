# Minly — URL Shortener

A small full-stack project built in 2022 to practice end-to-end application development: backend API design, database integration, frontend wiring, and automated deployment.

**Live:** [mly.vercel.app](https://mly.vercel.app)

---

## Overview

Minly is a URL shortener. You give it a long URL, it returns a short one. Visiting the short URL redirects to the original.

### Why build this?

At the time, I wanted to:

- Write a backend in Go without reaching for a heavy framework
- Work through the full cycle: design → build → deploy → maintain
- Set up CI/CD pipelines that deploy on release, without manual steps

This project is intentionally small in scope. It's not trying to demonstrate distributed systems, caching layers, or high-availability patterns. It's a clean, working service that covers the fundamentals.

---

## Design & Architecture

A few decisions worth noting:

- **No framework on the backend.** The HTTP server uses Go's standard library. Routing, middleware, and request handling are explicit — no magic.

- **Interface-driven design.** The `UrlStore` interface abstracts storage. The MongoDB implementation sits behind it. Swapping storage layers doesn't touch business logic.

- **ID generation with Indigo.** Short codes are generated using `osamingo/indigo`, a Sonyflake-based generator. This avoids database-level unique ID collisions without coordination overhead.

- **Validation before persistence.** URLs are parsed and DNS-checked before storing. Invalid or unreachable URLs fail fast.

- **Structured logging.** `zap` is used for logging, with environment-aware configuration (development vs production).

- **Graceful shutdown.** The server listens for OS signals and shuts down cleanly.

The frontend is a minimal Next.js app, included to complete the end-to-end flow rather than to showcase UI complexity.

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| Backend | Go (standard library) |
| Database | MongoDB |
| Frontend | Next.js |
| CI/CD | GitHub Actions |
| Deployment | Vercel (frontend), Docker-ready backend |

---

## Project Structure

```
backend/
├── main.go          # Entrypoint, wiring
├── config.go        # Environment and server configuration
├── api/             # Service interfaces and business logic
│   └── mongodb/     # Storage implementation
├── endpoints/       # HTTP handlers
└── server/          # HTTP server setup
```

---

## Running Locally

```bash
cd backend

export DB_URL="mongodb://..."
export DB_NAME="minly"

go run .
```

The server starts on port `8080` by default. See `config.go` for optional TLS and environment settings.

---

## CI/CD

GitHub Actions handle both continuous integration and deployment:

**On every push/PR to master:**
- Go backend: build and run tests
- Next.js frontend: build across Node 14.x and 16.x
- CodeQL security scanning for both Go and JavaScript

**On GitHub Release:**
- Backend: Docker image built and pushed to Heroku Container Registry, then released
- Frontend: Deployed to Vercel production

This separates CI from CD intentionally — merging to master doesn't auto-deploy. Deployments happen when a release is explicitly published, giving control over what goes live.

---

## Scope & Limitations

This project is a learning exercise, not a production system. It does not attempt to solve:

- Rate limiting or abuse prevention
- Analytics or click tracking
- Custom short codes or user accounts
- Horizontal scaling or caching

These are valid concerns for a real URL shortener, but they weren't the goal here. Later projects in my portfolio explore deeper backend architecture — database design, concurrency patterns, and system boundaries.

---

## License

[MIT](LICENSE)
