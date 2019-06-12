FROM golang:latest

ADD . /go/

WORKDIR /go/

RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux

CMD ["go", "run", "main.go"]
