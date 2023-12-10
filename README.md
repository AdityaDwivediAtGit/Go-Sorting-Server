# Go Sorting Server

This Go application serves as a demonstration of sequential and concurrent sorting of arrays. It provides two endpoints, `/process-single` and `/process-concurrent`, to showcase sorting methods with performance measurements.

## Features

- **Sequential Sorting**: The `/process-single` endpoint sorts each sub-array sequentially.
- **Concurrent Sorting**: The `/process-concurrent` endpoint sorts each sub-array concurrently using Go's concurrency features (goroutines, channels).
- **Performance Measurement**: Both endpoints measure the time taken to sort all sub-arrays and return the results in nanoseconds.

## Getting Started

### Prerequisites

- Go installed on your machine
- Docker installed for containerization

### Running the Server Locally

1. Clone this repository:

   ```bash
   gh repo clone AdityaDwivediAtGit/Go-Sorting-Server
   ```

2. Navigate to the project directory:

   ```bash
   cd go-sorting-server
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

   The server will be accessible at `http://localhost:8000`.

### API Endpoints

- **Sequential Processing:**

  ```bash
  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}' \
    http://localhost:8000/process-single
  ```

- **Concurrent Processing:**

  ```bash
  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}' \
    http://localhost:8000/process-concurrent
  ```

### Dockerization

This application can be containerized using Docker. To build and run the Docker image:

```bash
docker build -t go-sorter .
docker run -p 8000:8000 go-sorter
```

Or you can directly pull my Docker image and run the container:

```bash
docker pull infiniteintegrator/go-sorter
docker run -p 8000:8000 infiniteintegrator/go-sorter
```

### Accessing Railway Render Deployment

You can make requests to the deployed application on Railway Render using the provided URL:

```bash
# Example for sequential processing
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"to_sort": [[1, 22, 13], [4, 15, 6], [17, 8, 9]]}' \
  https://go-sorting-server-production-a169.up.railway.app/process-single
```

```bash
# Example for concurrent processing
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"to_sort": [[1, 22, 13], [4, 15, 6], [17, 8, 9]]}' \
  https://go-sorting-server-production-a169.up.railway.app/process-concurrent
```

### Example Run

![image](https://github.com/AdityaDwivediAtGit/Go-Sorting-Server/assets/107645490/abf3d419-b2a8-4485-8ff5-db74ac598f75)
