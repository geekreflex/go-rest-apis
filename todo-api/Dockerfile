FROM golang:1.20

# Set the working dir to /app
WORKDIR /app

# Copy the current dir content into the container at /app
COPY . .

# Download and install any required depencies
RUN go mod download

# Build the GO application
RUN go build -o todo-api

# Expose port 8000 for the application
EXPOSE 8000

# Set the evnironment variables for the MySQL database
ENV MYSQL_DSN "user:password@tcp(database:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"

# Run the Go application
CMD ["./todo-api"]