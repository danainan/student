FROM golang:1.21

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/student-server

EXPOSE 8000

CMD ["/app/student-server"]
