FROM golang:1.15.8-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0

# get dependencies for caching
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/captain cmd/main.go

FROM scratch AS bin
COPY --from=build /out/captain /
