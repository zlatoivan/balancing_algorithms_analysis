FROM golang:1.22
LABEL authors="ivan"

COPY go.mod go.sum ./
RUN go mod download

#RUN go install github.com/pressly/goose/cmd/goose

COPY . .

RUN go build -o main cmd/server/main.go

CMD ["./main"]



#WORKDIR /Homework
#RUN go get github.com/pressly/goose/cmd/goose@latest
#RUN go get -d -v ./...
#RUN go install -v ./...
