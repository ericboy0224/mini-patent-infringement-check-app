FROM golang:1.22.3-alpine
WORKDIR /app

# First copy only go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Create data directory
RUN mkdir -p /app/data

# Copy the data files first
COPY data/*.json /app/data/

# Then copy the rest of the application
COPY . .

RUN go build -o app
EXPOSE 8080
CMD ["./app"]