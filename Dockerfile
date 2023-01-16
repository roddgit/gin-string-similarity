FROM golang:1.18-alpine3.16

ENV GIN_MODE=release


RUN apk update && apk add --no-cache git && apk add bash && apk --no-cache add tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

# ENTRYPOINT ["/app/binary"]

# docker run --rm -i -t alpine /bin/sh --login