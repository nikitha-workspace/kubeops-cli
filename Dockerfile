# ---------- Build Stage ----------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o kubeops

# ---------- Runtime Stage ----------
FROM ubuntu:24.04

WORKDIR /app

COPY --from=builder /app/kubeops .

ENTRYPOINT ["./kubeops"]