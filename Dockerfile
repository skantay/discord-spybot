FROM golang:alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download


FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/spybot ./cmd/spybot

FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /bin/spybot /app
CMD ["/app"]