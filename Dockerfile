FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .


FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main .
COPY . .
EXPOSE 80
CMD ["/main"]