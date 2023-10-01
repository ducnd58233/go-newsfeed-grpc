FROM golang:1.20 AS builder

ARG DIR=/app

ENV GO111MODULE=on
WORKDIR $DIR

RUN mkdir /var/log/entry/
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/cache go mod download

RUN --mount=type=cache,target=/go/pkg/mod/cache CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app.linux -mod=mod cmd/main.go

RUN chmod +x /app/app.linux


FROM builder AS test_env

RUN cp /app/configs/files/test.yml /app/config.yml
CMD ["/app/app.linux"]


FROM builder AS live_env

RUN cp /app/configs/files/live.yml /app/config.yml
CMD ["/app/app.linux"]