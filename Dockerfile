FROM golang:1.20.5-alpine as builder

WORKDIR /app

COPY internal ./internal \
     cmd ./cmd \
     go.mod ./ \
     go.sum ./

RUN go mod download
RUN go build -o entrypoint

FROM alpine:3.18.2
COPY --from=builder /app/entrypoint /app/entrypoint

WORKDIR /app

ENTRYPOINT ["./entrypoint"]
