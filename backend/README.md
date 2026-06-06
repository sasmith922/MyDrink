# DrakeMaye Backend

Go + Gin REST API for DrakeMaye MVP.

## Run

```bash
cd backend
go mod tidy
go run ./cmd/server
```

Server runs on `http://localhost:8080`.

## Implemented Endpoints

- `GET /health`
- `GET /drinks`
- `GET /drinks/:id`
- `POST /drinks`
- `GET /logs`
- `POST /logs`
- `GET /logs/:id`
- `GET /feed`
- `POST /feed/:id/like`
- `GET /stats`
- `GET /stats/weekly`
- `GET /stats/monthly`
- `GET /profile`
- `PUT /profile/goals`

## Notes

- Uses SQLite with startup migrations + seed data.
- Uses a mock user (`id=1`) for MVP simplicity.
- Standard drink formula is implemented in `internal/utils/calculations.go`.
- CORS is enabled for Expo local development.

## Intentionally Not Implemented

- AI image recognition
- Image upload
- Computer vision processing
- Authentication
- Production deployment setup
