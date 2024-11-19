FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY lb/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /golb

EXPOSE 8080

CMD ["/golb"]