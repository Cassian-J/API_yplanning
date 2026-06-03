FROM golang:1.26 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiyplanning ./main.go

FROM scratch
COPY --from=builder /app/apiyplanning /apiyplanning

EXPOSE 8080
ENTRYPOINT ["/apiyplanning"]
