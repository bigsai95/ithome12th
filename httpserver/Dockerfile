# dockerfile
# Author: 阿榜
# Version: 1.0.0
FROM golang:1.15.2-alpine3.12
LABEL stage=httpserver
ENV PROJECT /go/src/httpserver
RUN cd /go/src
RUN mkdir -p $PROJECT
COPY ./ $PROJECT	
RUN cd $PROJECT && export GO111MODULE=on && go build -mod vendor 	
WORKDIR $PROJECT
CMD ["./httpserver"]