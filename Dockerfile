FROM golang:1.22

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLE=0 go build -o /usr/local/bin/foobar .

CMD [ "bash" ]
