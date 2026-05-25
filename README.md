# CNPS

**CNPS (Custom Network Protocol System)** is a lightweight, TCP-based, CLI-first network protocol designed for real-time, command-driven communication between clients, servers, and AI agents.

It focuses on:

* Minimal overhead
* Stream-based communication (NDJSON over TCP)
* Middleware-driven architecture
* Built-in authentication system (Cebu / Philippines model)
* Extensibility via command namespaces

---

## 🚀 Core Idea

CNPS replaces traditional request/response-heavy APIs (like REST/GraphQL) with a **persistent command stream model**.

Instead of:

```
HTTP request → response → close connection
```

CNPS uses:

```
Persistent TCP stream → NDJSON messages → continuous interaction
```

---

## 🧱 Architecture Overview

### Layers

* **Transport Layer**

    * TCP persistent connection
    * NDJSON framing (1 JSON object per line)

* **Protocol Layer**

    * Command routing (`cmd`)
    * Action semantics (`act`)
    * Target routing (`to`)
    * Metadata (`info`)
    * Payload (`payload`)

* **Application Layer**

    * Handlers
    * Middleware pipeline
    * Custom services (auth, logging, routing, etc.)

---

## 📦 Message Format

### Request

```json
{
  "type" : "REQ",
  "to": "service.name",
  "cmd": ".search.web",
  "act": "@APP.submit",
  "info": {
    "cebu" : {
      "philippines" : "auth.provider",
      "kid" : "uuid v7 key id",
      "token" : "Cebu CEBUTOKEN"
    },
    "request_id": "uuid"
  },
  "payload": {
    "query": "hello"
  }
}
```

### Response

```json
{
  "type": "RES",
  "status": 10,
  "info": {
    "request_id": "request uuid v7"
  },
  "payload": {
    "result": "ok"
  }
}
```

---

## ⚙️ Command System

CNPS uses a **dot-based capability tree**:

```
.search.web
.search.image
.auth.login
.system.ping
```

Each command maps to a handler:

```go
handler.AddRoutes(".test", testHandler)
```

---

## 🧩 Middleware System

CNPS supports chainable middleware:

```go
type Handler func(req dto.Request) (dto.Response, error)
type Middleware func(next Handler) Handler
```

Example:

```go
func Logger(next Handler) Handler {
    return func(req dto.Request) (dto.Response, error) {
        log.Println("Middleware triggered")
        return next(req)
    }
}
```

---

## 🔐 Authentication Model (Cebu / Philippines)

CNPS includes a custom auth system:

### Philippines (Auth Provider)

* Issues tokens
* Manages key lifecycle
* Handles multiservice identity

### Cebu (Token)

* Signed identity token
* Used across all CNPS services
* Enables SSO-like behaviour across systems

Key feature:

> One token → multiple services → unified identity layer

---

## 🌐 Design Goals

* Low latency (persistent TCP connection)
* CLI-first / AI-agent friendly
* No HTTP overhead
* Extensible command routing
* Middleware-based logic injection
* Distributed-friendly design (future-ready)

---

## 🔄 Data Flow

```
Client
  ↓
TCP Stream
  ↓
Middleware Chain
  ↓
Command Router
  ↓
Handler
  ↓
Response Stream
```

---

## 🛠 Example Server Setup

```go
server := cnps.NewServer(":31415")

server.Use(LoggerMiddleware)
server.AddRoutes(".test", testHandler)

server.Start()
```

---

## 📡 Why CNPS Exists

CNPS is designed for systems where:

* HTTP feels too heavy
* WebSocket feels too UI-oriented
* gRPC feels too rigid or schema-heavy

Instead, CNPS aims for:

> “terminal-native distributed systems”

---

## 🧪 Future Ideas

* `.cnp` executable workflow files (agent automation)
* Built-in event streaming model
* Distributed router nodes
* Pluggable auth providers (Cebu-compatible)
* AI-agent execution layer
* Pub/Sub command channels


