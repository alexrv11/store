
#build stage
FROM golang:alpine
WORKDIR /go/src/store
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go install /go/src/store
RUN ls /go/bin

ENTRYPOINT /go/bin/store

EXPOSE 1323
