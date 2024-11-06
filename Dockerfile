FROM node:20-alpine AS frontend

WORKDIR /frontend

# Install pnpm
RUN corepack enable && corepack prepare pnpm@latest --activate

# Copy frontend files
COPY frontend/package.json frontend/pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the frontend application
COPY frontend .

# Build frontend
RUN pnpm build

FROM golang:1.22.3-alpine AS backend

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Create required directories
RUN mkdir -p /app/data /app/prompts /app/frontend/dist

# Copy static files
COPY data/*.json /app/data/
COPY prompts/*.prompt /app/prompts/
COPY .env ./

# Copy source code
COPY . .

# Copy frontend build from frontend stage
COPY --from=frontend /frontend/dist /app/frontend/dist

# Build the application
RUN go build -o app

# Set environment variables from .env file
ENV $(cat .env | xargs)

# Set default MongoDB environment variables if not provided in .env
ENV MONGODB_URI=${MONGODB_URI:-mongodb://mongodb:27017}

EXPOSE 8080
CMD ["./app"]