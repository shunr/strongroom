FROM golang:1.14-buster

RUN apt-get update -y && apt-get install -y libpq-dev postgresql-client unzip

WORKDIR /tmp

# Install protoc
RUN PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

WORKDIR /strongroom_server

COPY go.mod .
COPY go.sum .

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go mod download
COPY . ./

RUN ./build.sh

ENV SERVER_PORT 13222
EXPOSE $SERVER_PORT
EXPOSE 13222
CMD ./start.sh