FROM golang:1.20

WORKDIR /home/go/go-whatsapp-rest

COPY . .

RUN cd .. && go mod tidy

RUN cd cmd && go build -o ../bin

ENTRYPOINT ["/home/go/go-whatsapp-rest/bin"]
