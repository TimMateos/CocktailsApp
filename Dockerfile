# --- Этап 1: Сборка (Builder) ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/CocktailApp main.go

# --- Этап 2: Финальный образ ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/CocktailApp .

COPY templates ./templates
COPY uploads ./uploads

EXPOSE 8080

CMD ["./CocktailApp"]