FROM golang:latest

LABEL maintener = "Izzat <arramsyah@gmail.com>"

WORKDIR /app

COPY go.mod . 

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build 

CMD ["./ProjectLogin-GOLANG"]