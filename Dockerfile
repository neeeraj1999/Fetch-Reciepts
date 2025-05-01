# Use Golang Alpine as base
FROM golang:1.20-alpine

# Define the working directory
WORKDIR /app

# Add go modules files
COPY go.mod go.sum ./

# Fetch dependencies
RUN go mod tidy

# Copy the rest of the code
COPY . .

# Compile the app
RUN go build -o app .

# Open port 8080
EXPOSE 8080

# Execute the binary
CMD ["./app"]
