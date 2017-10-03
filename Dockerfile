FROM golang:alpine as builder
COPY src /go/src/docker-worker
WORKDIR /go/src/docker-worker
RUN apk add --no-cache git gcc libc-dev \
    && go get
RUN GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -installsuffix netgo -ldflags '-w' -o main .

FROM scratch
ADD Go_Daddy_Class_2_CA.pem /etc/ssl/certs/
COPY --from=builder /go/src/docker-worker/main .
CMD ["./main"]
