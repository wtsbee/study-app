FROM golang:1.19.1-alpine

WORKDIR /go/src

RUN apk update && apk add git

# sql-migrateの依存パッケージ
RUN apk add --no-cache gcc && apk add --no-cache musl-dev

# ビルドによく使われるものを集めたパッケージをインストールする場合
# RUN apk add build-base

# sql-migrateのインストール
RUN go install github.com/rubenv/sql-migrate/...@latest
