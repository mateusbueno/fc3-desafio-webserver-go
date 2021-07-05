FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN GOOS=linux go build main.go

EXPOSE 8000

CMD ["./main"]
