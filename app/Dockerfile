FROM golang:1.13-alpine3.10

WORKDIR /go/angles/
ADD . .

RUN go get -d -v ./... && \
    go build angles

#####################################################################
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /opt/app/

COPY --from=0 /go/angles/angles /opt/app/angles
COPY templates/ /opt/app/templates/

ENTRYPOINT ["./angles"]