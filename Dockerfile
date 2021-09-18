FROM golang:1.17.1-alpine
RUN apk update apk add git

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod .
COPY go.sum .

RUN go mod tidy
COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]
