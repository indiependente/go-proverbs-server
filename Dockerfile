FROM golang:alpine as builder

WORKDIR $GOPATH/src/github.com/indiependente/go-proverbs-server/
ADD . .
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /service

# Deployable image
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /service /app/
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/app/service"]
