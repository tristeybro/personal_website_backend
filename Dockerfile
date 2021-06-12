FROM golang:1.15-alpine

ARG SENDGRID_API_KEY="not-a-real-api-key"

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod vendor
RUN go mod verify
ENV SENDGRID_API_KEY=$SENDGRID_API_KEY
RUN go build -o main
RUN ls .

CMD ./main
