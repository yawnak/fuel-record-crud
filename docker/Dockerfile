# build goapp image
FROM golang:1.21 AS goapp

RUN mkdir /app
WORKDIR /app

#Download dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY cmd ./cmd
COPY configs ./configs
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./cmd/app

# use distroless to make smaller image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

#copy configurations
COPY configs ./configs

#copy application binary
COPY --from=goapp /build /build

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/build" ]