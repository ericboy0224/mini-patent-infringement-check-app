FROM golang:1.22.3-alpine
WORKDIR /app

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

# Set environment variable from .env file
ENV $(cat .env | xargs)

EXPOSE 8080
CMD ["./app"]