# CLAUDE.md

## How to work with me

**Do not write or edit code in this repo unless I explicitly ask you to.** This is a learning project. I am teaching myself Go and backend design by building it, so the default is: I write the code, you explain.

- Default mode is answer and explain. When I ask "what's next" or "how do I do X", tell me what to do and why. Do not open an editor and do it for me.
- Explain the real mechanism one layer deep. Say what actually happens (headers flush on first body write, interfaces let me swap implementations, etc.), not a surface-level recipe and not an exhaustive dump.
- Short code snippets in your reply are fine when they clarify a mechanic. Writing them into my files is not, unless I say so.
- Exceptions: when I directly tell you to create or change a specific file (like this one), do it. Reading files, running `go build`, `go vet`, `go test`, and git read commands to understand state is always fine.
- If I ask you to code something, do exactly that scope and stop. Don't refactor neighboring code.

## Project

AI degree planner for Georgia Tech students. Input: courses already taken plus chosen CS thread(s). Output: a multi-semester (up to 8) plan that satisfies prereqs, credit caps, and term availability. Go backend. Inspired by a BU hackathon project; chosen over a Raft project.

### Layout (planned)

- Module: `github.com/imermigkas3/scheduler`
- `cmd/api` — HTTP server (exists, one hello-world route on :3000)
- `cmd/ingest` — run-once worker, idempotent upserts into Postgres
- `internal/planner` — prereq graph, scheduling engine
- `internal/ingest` — parsing GT Scheduler crawler per-term JSON
- Docker Compose for local Postgres
- Data source: fork of the GT Scheduler crawler's per-term JSON

### Roadmap (set Oct 2026)

1. Setup + data: repo, schema (courses, prereqs, term availability), ingest worker.
2. Planner engine: prereq graph with AND/OR, Kahn's topological sort, credit caps, term availability, critical-path priority, verifier, unit tests, hand-encoded thread requirements as JSON.
3. API (list courses, generate/save plan) + simple React page + Docker deploy + GitHub Actions.
4. AI layer: natural language to structured constraints, validated in Go; AI-generated plan explanations.
5. Optional: Python CP-SAT service, with the Go heuristic as fallback.
6. Polish: caching, rate limiting, more majors, Course Critique GPA data.

### Planned API surface (phase 3)

- `GET /healthz`
- `GET /courses?term=&subject=`
- `GET /courses/{id}`
- `POST /plans` (body: courses taken + thread) returns a plan
- `GET /plans/{id}`, `POST /plans/{id}/save` once plans persist

### Conventions

- Standard library `net/http` with Go 1.22+ `ServeMux` patterns (method prefix, `{id}` path values, `{$}` for exact root).
- Handlers depend on small interfaces (e.g. a course store) so Postgres can replace an in-memory stub without touching handlers.
- Prefer backend/infra depth over frontend polish.

## About me

GT CS student (took CS 2200). Learning Go. Prefer backend and infra work over frontend.
