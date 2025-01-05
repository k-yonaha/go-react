# Go × React 競艇チャット

- 24場の競艇場のチャット

## 構成
Docker 27.4.0
Golang 1.23.3-alpine
React 18.3.17

## ローカル環境構築手順

#### .env作成

```
$ touch backend/.env
```

#### .env編集する(仮で値を入れています)
```
DB_USER=db-user
DB_PASSWORD=db-pass
DB_NAME=sample
DB_HOST=postgres
DB_PORT=5432
DBMS=postgres
```
#### docker起動

```
$ docker compose up -d
```
#### dockerに入る

```
$ docker compose exec go /bin/sh
```

#### 各競艇場の部屋を作成するseeder実行

```
$ go run seeder/seeder.go
```

#### 本日のレース情報の取得とDB保存

```
http://localhost:8080/api/downloadSchedule?date=[日付(YYYYMMDD)]
```