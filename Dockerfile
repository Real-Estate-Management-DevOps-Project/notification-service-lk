FROM golang:1.24-bookworm AS build
WORKDIR /src

# Copy dependency files first to leverage build cache
COPY go.mod go.sum ./

# Update system packages, install git only for downloading modules, then remove it
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git && \
    go mod download && \
    apt-get remove -y git && apt-get autoremove -y && apt-get clean

# Copy the rest of the source and build a statically linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w' -o /notification-service ./

FROM scratch
COPY --from=build /notification-service /notification-service
EXPOSE 3000
ENTRYPOINT ["/notification-service"]
