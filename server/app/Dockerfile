# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23.4 AS build-stage

WORKDIR /app

# TODO: consider using swagger cli to generate the swagger API documentation code
COPY src/go/ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /software-slayer-server

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /software-slayer-server /software-slayer-server

USER nonroot:nonroot

ENTRYPOINT ["/software-slayer-server"]