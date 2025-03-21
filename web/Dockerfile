FROM golang:1.23

WORKDIR /gym_tracker

COPY go.mod go.sum /gym_tracker/

RUN go mod download

COPY . .

EXPOSE 3000
