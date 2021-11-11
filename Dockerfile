FROM golang:1.16-alpine

ENV GOPATH /go

WORKDIR /go/src/stone-test

RUN apk add build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY api/ ./api
COPY middlewares/ ./middlewares
COPY models/ ./models
COPY services/ ./services
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]