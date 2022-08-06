FROM golang:1.17

RUN apk update && apk add --no-cache git

WORKDIR /home/go/go-whatsapp-rest

COPY . .

RUN go mod tidy

RUN cd cmd && go build -o ../bin

ENTRYPOINT ["/home/go/go-whatsapp-rest/bin"]
