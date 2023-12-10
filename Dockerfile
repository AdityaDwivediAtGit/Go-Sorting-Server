FROM golang:latest
WORKDIR /app

# Copy the local package files to the container's working directory
COPY . .

RUN go build -o main .
EXPOSE 8000
CMD ["./main"]