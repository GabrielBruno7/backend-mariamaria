FROM golang:1.22.2

WORKDIR /go/src/app

COPY . . /go/src/app/

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD [ "./main" ]