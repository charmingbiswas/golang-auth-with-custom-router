# Multi-stage Dockerfile

# Build stage
FROM golang:1.21.4-alpine3.17 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/main .


EXPOSE 4000
CMD [ "/app/main" ]