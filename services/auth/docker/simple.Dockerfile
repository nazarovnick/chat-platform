FROM golang:1.24.0-alpine3.21
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bin/main ./cmd/server
EXPOSE 8080

CMD ["/bin/main"]