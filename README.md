# ID Service

A lightweight Go REST API that generates an `eduID` for a given email address.

## What it does

Accepts a `GET` request with an `email` query parameter and returns a randomly generated `eduID` in the format `<email>+<6-digit-number>` (e.g. `alice@bar.com+871715`).

## Project structure

```
.
├── main.go     # Entire service — handler, models, business logic
├── go.mod      # Go module definition
└── README.md
```

---

## Running locally

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) — verify with `go version`

### Steps

```bash
# 1. Enter the project directory
cd id-service

# 2. Tidy dependencies
go mod tidy

# 3. Run
go run main.go
```

You should see:
```
EduID Service starting on :8081
```

### Test with curl

**Happy path**
```bash
curl "http://localhost:8081/user-attributes/edu-id?email=alice@bar.com"
```
```json
{ "eduID": "alice@bar.com+871715" }
```

**Missing email parameter**
```bash
curl "http://localhost:8081/user-attributes/edu-id"
```
```json
{ "error": "missing_parameter", "description": "Query parameter 'email' is required" }
```

---

## API reference

### `GET /user-attributes/edu-id`

**Query parameters**

| Parameter | Required | Description |
|---|---|---|
| `email` | ✅ Yes | The user's email address |

**Responses**

| HTTP | When | Body |
|---|---|---|
| 200 | Success | `{ "eduID": "<email>+<6-digit-number>" }` |
| 400 | `email` param missing | `{ "error": "missing_parameter", "description": "..." }` |
| 405 | Non-GET request | `{ "error": "method_not_allowed", "description": "..." }` |
