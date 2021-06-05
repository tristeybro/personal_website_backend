FROM golang:1.15-alpine

ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go build -o main

CMD ["./personal_website_backend"]
