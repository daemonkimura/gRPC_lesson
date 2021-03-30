FROM golang:1.14.4-buster

WORKDIR /go/src/app

#アップデートとGitのインストール
RUN apt-get update && \
    apt-get install -y git zip

#protocol bufferのインストール
WORKDIR /tmp
ENV PROTOBUF_URL https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip

RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
RUN ls /tmp
RUN unzip /tmp/protobuf.tar.gz
RUN mv bin/* /usr/local/bin/ && \
    mv include/* /usr/local/include/

RUN protoc --version
RUN rm -rf *

#Go言語関連
WORKDIR /go/src/app

#gRPCのインストール
RUN go get -u google.golang.org/grpc
#protocol bufferプラグインのインストール
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

RUN go mod init grpcTest
#RUN go get -v -u github.com/gin-gonic/gin
#RUN go get -v github.com/jinzhu/gorm
#RUN go get -v github.com/go-sql-driver/mysql
#RUN go get -v github.com/dgrijalva/jwt-go
#RUN go get -v github.com/joho/godotenv
#RUN go get -v github.com/gin-contrib/cors
#RUN go get -v -u gopkg.in/src-d/go-git.v4/...
#RUN go get -v -u gopkg.in/src-d/go-billy.v4/...



