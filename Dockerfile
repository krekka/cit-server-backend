FROM golang:1.20
WORKDIR /app

# Downloading dependencies
COPY go.mod go.sum ./
RUN go mod download

# Building application
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o build/api

# Running our application
EXPOSE 8090

ENTRYPOINT [ "build/api" ]
CMD [ "serve", "--dir", "/pb_data", "--http", "0.0.0.0:8090" ]