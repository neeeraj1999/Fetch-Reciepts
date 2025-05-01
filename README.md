# Receipt Processor

## Overview
A RESTful web service built in Go that processes retail receipts and calculates reward points according to predefined rules. The service can be deployed either using Docker or run locally.

## Features
- REST API endpoint for submitting and processing receipts
- Automated points calculation based on receipt details
- In-memory storage system for receipt data

## Prerequisites
- Docker (for containerized deployment)
- Go (for local development)

## Installation

### Using Docker
1. Clone the repository:
   ```bash
   git clone https://github.com/neeeraj1999/Fetch-Reciepts.git
   cd Fetch-Reciepts
   ```

2. Build the Docker image:
   ```bash
   docker build -t Fetch-Receipts .
   ```

3. Start the container:
   ```bash
   docker run -d -p 8080:8080 Fetch-Reciepts
   ```

### Local Development
1. Clone the repository:
   ```bash
   git clone https://github.com/neeeraj1999/Fetch-Reciepts.git
   cd Fetch-Reciepts
   ```

2. Run the application:
   ```bash
   go run main.go
   ```
