# syntax=docker/dockerfile:1
# single stage build \/
FROM golang:1.22-alpine

ENV APP_PORT=8000

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /ticketsystem

EXPOSE $APP_PORT
CMD ["/ticketsystem"]

# multi-stage build \/
# FROM golang:1.22-alpine AS build-stage

# ENV APP_PORT=8000

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /ticketsystem

# # Run the tests in the container
# FROM build-stage AS run-test-stage

# RUN go test -v ./...

# # Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage


# WORKDIR /

# COPY --from=build-stage /ticketsystem /ticketsystem

# EXPOSE $APP_PORT

# USER nonroot:nonroot

# ENTRYPOINT ["/ticketsystem"]