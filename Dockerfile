FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build -v -o imageapi ./...

EXPOSE 8080
CMD ["/app/imageapi"]