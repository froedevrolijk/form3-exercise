FROM golang:1.18
 
WORKDIR /form3
 
COPY go.mod go.sum ./

RUN go mod download

COPY . .