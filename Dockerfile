FROM golang:1.16 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 9091
CMD ["./main"]
