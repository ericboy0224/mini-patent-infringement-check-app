FROM golang:1.22.3-alpine
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# First copy only go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Create required directories
RUN mkdir -p /app/data /app/prompts

# Copy static files
COPY data/*.json /app/data/
COPY prompts/*.prompt /app/prompts/
COPY .env ./

# Copy source code
COPY . .

# Build the application
RUN go build -o app

# Set environment variables from .env file
ENV $(cat .env | xargs)

# Set default MongoDB environment variables if not provided in .env
ENV MONGODB_URI=${MONGODB_URI:-mongodb://mongodb:27017}

EXPOSE 8080
CMD ["./app"]