FROM golang:1.20.6-alpine3.18

WORKDIR /go-sosmed

COPY go.mod /go-sosmed

RUN go mod download

COPY . /go-sosmed

RUN  go mod tidy

RUN go build -o /go-sosmed/bin/main /go-sosmed/cmd/main.go

EXPOSE 50051

CMD [ "/go-sosmed/bin/main" ]