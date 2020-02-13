FROM golang:alpine as builder
RUN mkdir /build 
ADD ./data-generator /build/
WORKDIR /build
RUN apk update
RUN apk add git
RUN go get github.com/gocql/gocql
RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]