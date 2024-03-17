FROM golang:alpine as builder
WORKDIR /go/src
COPY . .
RUN go get
RUN go build -o /go/bin/app_api

FROM alpine
RUN apk update && apk add --no-cache tzdata
COPY --from=builder /go/bin/app_api /app_api
ENV TZ="Asia/Bangkok"
ENTRYPOINT [ "/app_api" ]