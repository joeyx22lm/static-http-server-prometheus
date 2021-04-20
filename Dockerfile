FROM golang:1.16.3-alpine3.13

ENV PROMETHEUS_PORT 8080
ENV PORT 5000

WORKDIR /usr/src/app

COPY ./src .
RUN go install
ADD ./www /www

CMD ["go", "run", "main.go"]
