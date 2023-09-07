FROM golang:1.20.2-alpine3.17 as builder

RUN apk add --no-cache git alpine-sdk

ADD . /go/src/workspace
WORKDIR /go/src/workspace

RUN make build

# use a minimal alpine image
FROM alpine:3.17

# set working directory
WORKDIR /go/bin

COPY --from=builder /go/src/workspace/main .

RUN touch error.log
RUN chmod 777 error.log

USER 1001
# run the binary
CMD ["./main"]
