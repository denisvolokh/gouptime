# GoUpTime Project – Specification

## Overview
GoUpTime is a multi-tenant monitoring platform inspired by updown.io.  
It supports monitoring multiple resource types, runs probes from different regions, stores time-series metrics, and delivers alerts via various channels.

---

## Features (v1)

- **Multi-Resource Monitoring**
    - HTTP(S): status code, latency, keyword match
    - TCP port: connectivity & handshake
    - ICMP ping: packet loss & RTT
    - DNS: resolve time & mismatch detection
    - SMTP/STARTTLS: banner and TLS checks
    - SSL/TLS expiry: certificate validity & chain
    - Heartbeat: “ping us within N seconds”

- **Multi-Tenant Model**
    - Organizations → Projects → Resources → Checks
    - RBAC: owner, member roles
    - API keys per project, scoped

- **Checks**
    - Configurable interval, timeout, expectations
    - Multi-region probing with quorum logic
    - Maintenance windows to suppress alerts

- **Execution Flow**
    1. Scheduler emits due jobs
    2. Agents run probes (per region)
    3. Results sent to message bus
    4. Ingestor validates & stores results
    5. Alert engine evaluates state transitions
    6. Notifier delivers alerts

- **Data & Incidents**
    - Store results in Postgres/Timescale
    - Track incidents with FSM (OK → WARN → DOWN)
    - SLA & uptime reporting
    - Public/private status pages

- **Notifications**
    - Email, Slack, SMS, webhook
    - Escalation policies (per project)

---

## Core Data Model

- `organizations` – tenant container
- `projects` – group resources & checks
- `resources` – endpoints or services to monitor
- `checks` – probe definition (kind, interval, config)
- `check_runs` – individual executions
- `incidents` – downtime or degraded states
- `notification_channels` – email, Slack, SMS, webhook
- `escalation_policies` – rules for alert escalation
- `maintenance_windows` – alert suppression periods

---

## MVP Scope

- Resource types: HTTP(S), ICMP, TLS expiry, Heartbeat
- Single region of agents
- Email & Slack notifications
- Basic incident tracking
- Public status page per project

---

## Extensibility

- New probes can be added by implementing the `Probe` interface in Go:
  ```go
  type Probe interface {
      Kind() string
      Validate(cfg json.RawMessage) error
      Execute(ctx context.Context, target Target, cfg json.RawMessage) (Result, error)
  }

# Tech Stack

## Backend

| Component         | Technology / Library                    |
|-------------------|-----------------------------------------|
| Language          | Go                                       |
| API Framework     | Gin (REST API)             |
| Messaging Bus     | NATS JetStream                           |
| Database          | PostgreSQL + TimescaleDB (time-series)   |
| Migrations        | golang-migrate                           |
| ORM/Driver        | pgx                                      |
| Metrics/Obs       | Prometheus client_golang, structured logs|
| CLI               | spf13/cobra                              |

---

## Services

- **API** – REST API (multi-tenant, auth, CRUD for resources & checks)
- **Scheduler** – emits check jobs at intervals with jitter
- **Agent** – stateless prober (HTTP, ICMP, DNS, etc.), runs per region
- **Ingestor** – validates & stores results, aggregates metrics
- **Alerter** – state machine for incidents & escalations
- **Notifier** – delivers alerts (email, Slack, SMS, webhook)

---

## Frontend

| Component      | Technology                |
|----------------|---------------------------|
| Web Dashboard  | HTMX + Alpine.js          |
| UI Styling     | TailwindCSS (if React)    |
| Status Pages   | Server-rendered (HTMX)    |

---

## Resource Probing Libraries

- **HTTP(S):** `net/http` (custom transport)
- **ICMP:** `golang.org/x/net/icmp`
- **DNS:** `miekg/dns`
- **TLS/SSL:** `crypto/tls`
- **SMTP/STARTTLS:** `net/smtp`, `crypto/tls`

---

## Deployment

| Environment    | Technology / Service           |
|----------------|--------------------------------|
| Dev            | Docker Compose                 |
| Prod           | Kubernetes / Helm              |
| Agents         | Fly.io, Hetzner, DigitalOcean   |
| Observability  | Prometheus + Grafana            |
| Secrets Mgmt   | HashiCorp Vault / KMS           |

---

## Security

- JWT / API keys per project
- Signed webhooks (HMAC)
- Encrypted secrets (Vault/KMS)
- Multi-tenant row-level access controls

---
