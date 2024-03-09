FROM golang:1.21

WORKDIR /home/go/go-whatsapp-rest

COPY . .

RUN go mod tidy

RUN cd cmd && go build -o ../bin

ENTRYPOINT ["/home/go/go-whatsapp-rest/bin"]
