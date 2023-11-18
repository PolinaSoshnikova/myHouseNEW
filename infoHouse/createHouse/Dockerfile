FROM ubuntu:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod dowload
RUN go build -o rudn./main.go

CMD["*./rudn"]