FROM golang:1.22

WORKDIR /app
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-world

EXPOSE 8080

# Run
CMD ["/hello-world"]
