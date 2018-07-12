FROM golang:alpine as builder 

ENV SRC=/go/src/github.com/ComputePractice2018/socialnetwork/backend/ 
RUN mkdir -p '${SRC}' 
ADD . ${SRC} 

RUN ls ${SRC} 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/socialnetwork github.com/ComputePractice2018/socialnetwork/backend 


FROM alpine:3.8 

WORKDIR /opt 
COPY -from=bulilder /go/bin/socialnetwork 


ENTRYPOINT ["/go/bin/socialnetwork"]