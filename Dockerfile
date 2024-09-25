# Base Stage for building Go binaries
FROM golang:1.23-alpine AS base


# Adding necessary tools like git, bash, and openssh
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

ENV CGO_ENABLED=0

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Optionally, run gqlgen generate here
RUN go run github.com/99designs/gqlgen generate

# Build the Main Go app
RUN go build -o main .

# Build the Seeder Binary
# RUN go build -o seed_db cmd/seed.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the main application
CMD ["./main"]
