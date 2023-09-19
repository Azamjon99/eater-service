FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --update tzdata

WORKDIR /root/

COPY ./bin/server /go/bin/app
WORKDIR /go/src/app

EXPOSE 3030

ENTRYPOINT ["/go/bin/app"] 