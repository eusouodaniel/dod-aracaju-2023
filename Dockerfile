FROM golang:latest as builder

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .

FROM scratch

COPY --from=builder /app/server .

CMD ["./server"]