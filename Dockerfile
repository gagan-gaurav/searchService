# Go base image
FROM golang:1.21.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /searchservice ./cmd/main.go

EXPOSE 8081

# Run
CMD ["/searchservice"]