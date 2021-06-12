FROM golang:1.15-alpine

ARG SENDGRID_API_KEY="not-a-real-api-key"

ENV GO111MODULE=on
WORKDIR /app
COPY . .
ENV SENDGRID_API_KEY=$SENDGRID_API_KEY
RUN go build -o main

CMD ["./main"]
