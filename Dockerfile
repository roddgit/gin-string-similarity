FROM golang:1.18-alpine3.16

ENV GIN_MODE=release

ENV ELASTICAPM_RUN=$GIN_ELASTICAPM_RUN
ENV ELASTICAPM_SERVICE_NAME=$GIN_ELASTICAPM_SERVICE_NAME
ENV ELASTICAPM_SERVER_URL=$GIN_ELASTICAPM_SERVER_URL
ENV ELASTICAPM_ENV=$GIN_ELASTICAPM_ENV


RUN apk update && apk add --no-cache git && apk add bash && apk --no-cache add tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

# ENTRYPOINT ["/app/binary"]

# docker run --rm -i -t alpine /bin/sh --login