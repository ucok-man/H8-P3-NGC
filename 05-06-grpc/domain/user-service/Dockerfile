# Build the application from source
FROM golang:1.21.6 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-service

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v -cover ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app-service /app-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app-service"]