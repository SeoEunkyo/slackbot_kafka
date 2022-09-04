FROM golang:latest
WORKDIR /go/github.com/seoEunkyo/slackbot_kafka
Copy . .
WORKDIR src/botlistener

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o botlistener

FROM alpine

COPY --from=0 /go/github.com/seoEunkyo/slackbot_kafka/src/botlistener /go/botlistener
WORKDIR go/botlistener
ENV LISTEN_URL=0.0.0.0:8282
EXPOSE 8282
CMD ["./botlistener"]
