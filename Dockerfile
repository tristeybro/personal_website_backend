FROM golang:1.15-alpine

# RUN apk --no-cache add ca-certificates

ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY . .
# RUN go mod download
# RUN go mod vendor
# RUN go mod verify

# RUN mkdir src
# COPY src/main.go src/main.go
# WORKDIR /
# RUN unset GOPATH
# RUN go mod download
# RUN go mod verify

RUN ls .
RUN go build -o main


# CMD ["go", "run", "main.go"]
CMD ["./personal_website_backend"]
