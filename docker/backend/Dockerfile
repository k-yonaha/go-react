FROM golang:1.23.3-alpine

WORKDIR /go/src/app

RUN apk --no-cache add git gcc musl-dev tzdata

COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

# Goのソースコードをコピー
COPY backend .

# airをインストール
RUN go install github.com/air-verse/air@latest

# airを使って開発サーバーを起動
CMD ["air"]
