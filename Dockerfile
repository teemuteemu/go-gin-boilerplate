FROM golang:latest AS build

# RUN mkdir /app
COPY . /app/

WORKDIR /app

RUN CGO_ENABLED=0 go test -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY --from=build /app/main .

ENV HTTP_PORT=8000
ENV DB_NAME=gorm_test
ENV DB_USER=user
ENV DB_PASSWORD=passu
EXPOSE 8000

CMD ["./main"]
