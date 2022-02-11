###########################
# Build container for app #
###########################

# Build arg VERSION only works before the FROM statement
ARG VERSION=latest
FROM golang:${VERSION} AS builder

# Sets build args for container
ARG GITHUB_OWNER
ARG GITHUB_REPO
ARG SERVICE_NAME

# Install task binary to run tasks
RUN  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# Set proper workdir for build
WORKDIR /go/src/github.com/${GITHUB_OWNER}/${GITHUB_REPO}

# Copy, get all packages and build application
COPY . .
RUN go mod tidy
RUN task build

##############################
# Deployment container build #
##############################

FROM alpine:latest

# Sets build args for container
ARG GITHUB_OWNER
ARG GITHUB_REPO
ARG SERVICE_NAME

# Install Root Certificates for https endpoints
RUN apk --no-cache add ca-certificates

# Set workdir for app and copy from the builder image
WORKDIR /root/
COPY --from=builder /go/src/github.com/${GITHUB_OWNER}/${GITHUB_REPO}/builds/$SERVICE_NAME-linux-amd64 ./app
CMD ["./app"]