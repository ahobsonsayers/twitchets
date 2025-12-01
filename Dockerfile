# Frontend builder Image
FROM node:24-slim AS frontend-builder

WORKDIR /app

# Setup pnpm
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

# Install dependencies
COPY ./frontend/package.json .
COPY ./frontend/pnpm-lock.yaml .
RUN pnpm install --frozen-lockfile

# Build
COPY ./frontend .
RUN pnpm run build

# Backend builder Image
FROM golang:1.24 AS backend-builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
COPY --from=frontend-builder /app/dist ./frontend/dist
RUN go build -v -o ./bin/ .

# Distribution Image
FROM alpine:latest


RUN apk add --no-cache libc6-compat

COPY --from=backend-builder /app/bin/twitchets /usr/bin/twitchets

EXPOSE 9000

WORKDIR /twitchets

ENTRYPOINT ["/usr/bin/twitchets"]
