FROM golang:1.16-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apt-get update
RUN apt-get -y install postgresql-client 

COPY . ./

RUN go build -o /my-app cmd/app/main.go

EXPOSE 3000

CMD [ "/my-app" ]

