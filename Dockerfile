FROM golang:1.19-bullseye

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main main.go

ENTRYPOINT ["./main"]
