FROM golang:1.22-alpine AS builder
WORKDIR /app

# Enable build kit cache mounts for faster re-builds
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux CGO_ENABLED=0 go build -o myapp .


FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]