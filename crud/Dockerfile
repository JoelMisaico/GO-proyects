FROM golang:1.18.4-alpine

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o crud-app

EXPOSE 3000

CMD [ "./crud-app" ]