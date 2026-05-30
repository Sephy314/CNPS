# CNPS

**CNPS (Command Network Protocol Stream)** is a command-oriented application protocol built on persistent TCP connections and NDJSON framing.

It is designed for CLI applications, AI agents, real-time systems, and lightweight services that require a simple stream-native communication model.

---

## Features

* Persistent TCP connections
* NDJSON framing
* Command-oriented architecture
* Request/Response model
* Bidirectional streaming
* Middleware support
* Context propagation
* CLI-first design
* AI Agent-friendly semantics
* Human-readable protocol

---

## Why CNPS?

CNPS focuses on **commands instead of resources**.

Instead of:

```http id="h1q9xk"
POST /user
```

or:

```proto id="k2j0aa"
rpc CreateUser()
```

CNPS uses:

```json id="c9p2lm"
{
  "cmd": ".user",
  "act": "@MAK"
}
```

This makes intent explicit and consistent across systems.

---

## Architecture

```text id="a8d1qs"
Application
     │
 Middleware
     │
   Router
     │
    CNPS
     │
 Persistent TCP
```

CNPS operates over long-lived TCP connections using NDJSON messages.

Each line is a complete message.

---

## Message Format

### Request

```json id="m4x9za"
{
  "to": "chat.app",
  "cmd": ".message",
  "act": "@MAK",
  "info": {},
  "payload": {
    "content": "Hello"
  }
}
```

### Response

```json id="p7v3ld"
{
  "status": 10,
  "info": {},
  "payload": {
    "message": "OK"
  }
}
```

---

## Fields

| Field   | Description          |
| ------- | -------------------- |
| to      | Target application   |
| cmd     | Command / capability |
| act     | Action type          |
| info    | Metadata             |
| payload | Data body            |

---

## Actions (CRUD Model)

| Action | Meaning      |
| ------ | ------------ |
| @QRY   | Query / Read |
| @MAK   | Create       |
| @UDT   | Update       |
| @RMV   | Delete       |

---

## Command Tree

```text id="t1x8ab"
.user
.user.profile
.user.settings

.post
.post.comment

.chat
.chat.room
.chat.message
```

---

## Status Codes

### 1. Success (SUC)

| Code | Meaning |
| ---- | ------- |
| 10   | OK      |

---

### 2. Processing

| Code | Meaning    |
| ---- | ---------- |
| 20   | Accepted   |
| 21   | Processing |

---

### 3. Client Error

| Code | Meaning           |
| ---- | ----------------- |
| 30   | Client Error      |
| 31   | Bad Request       |
| 32   | Not Authorised    |
| 33   | Expired Token     |
| 34   | Forbidden         |
| 35   | Duplicated        |
| 36   | Conflict          |
| 37   | Too Many Requests |
| 38   | Not Found         |

---

### 4. Server Error

| Code | Meaning               |
| ---- | --------------------- |
| 40   | Internal Server Error |

---

## Streaming

CNPS supports continuous bidirectional streaming over a single TCP connection.

```text id="s9k2qp"
Client → Request
Server → Response

Server → Event
Server → Event
Server → Event
```

---

## Middleware

Middleware provides extensibility for cross-cutting concerns.

Examples:

* Logging
* Authentication
* Recovery
* Metrics
* Rate Limiting
* Tracing

```go id="l2m8qz"
server.Use(Logger())
server.Use(Recovery())
server.Use(Auth())
```

---

## Routing

```go id="r7n1vd"
router.Handle(".user", HandleUser)
router.Handle(".post", HandlePost)
router.Handle(".chat", HandleChat)
```

---

## Use Cases

### CLI

```bash id="c1v9kk"
cnps call app .user @QRY
```

### AI Agents

```json id="q3p8xx"
{
  "cmd": ".calendar",
  "act": "@MAK",
  "payload": {
    "title": "Meeting"
  }
}
```

### Real-Time Systems

* Chat systems
* Multiplayer games
* Event streaming
* Notifications

### Internal Systems

* Microservices
* Workers
* Control planes
* Internal APIs

---

## Design Goals

* Simplicity over complexity
* Stream-first communication
* Human-readable protocol
* Middleware extensibility
* AI-friendly structure
* High-performance TCP communication

---

## Non-Goals

CNPS is not:

* A replacement for HTTP
* A message broker
* A service mesh
* A binary protocol

CNPS is an application-layer protocol over TCP.

---

## License

Licensed under the Apache License, Version 2.0.

See LICENSE file for details.
