FROM golang:1.18.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /my-server

EXPOSE 8080

CMD [ "/my-server" ]