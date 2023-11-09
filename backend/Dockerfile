FROM golang:1.21.4-alpine3.18

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN mkdir -p /usr/local/go/src/app/

WORKDIR /usr/local/go/src/app

COPY . .

RUN apk update \
  && apk add --no-cache git \
  && apk add git \
  && go install github.com/cosmtrek/air@v1.29.0 \
  && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]