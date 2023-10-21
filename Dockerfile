FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go */*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /kaktus

EXPOSE 8080

CMD [ "/kaktus" ]
