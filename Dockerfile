FROM golang:1.15-alpine

ENV GO111MODULE=on
WORKDIR /app
COPY . .
# TODO: Figure out a way to source credentials securely.
RUN source sendgrid.env
RUN go build -o main

CMD ["./main"]
