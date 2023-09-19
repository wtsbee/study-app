# study-app

## 概要

- カンバン方式の機能とマークダウンによるドキュメント作成機能を持つアプリです
- 作成したリストやカードをドラッグすると移動できます
- カードをクリックするとマークダウンの入力画面に遷移します
- WebSocket で通信しており、別画面で入力した内容がリアルタイムで反映されます

## 初回起動

### コンテナ起動

`golang`ディレクト配下に`.env`ファイルを作成し、`.env.sample`の中身をコピーして貼り付けます。

以下のコマンドを実行し、コンテナを起動します。

```
$ docker-compose up -d
```

### backend のコンテナ内での作業

以下のコマンドを実行します。

```
$ go mod tidy
```

サーバーを起動します。

```
$ go run main.go
```

### frontend のコンテナ内での作業

以下のコマンドを実行します。

```
$ npm install
```

サーバーを起動します。

```
$ npm run dev
```

## 2 回目以降のコンテナ起動

`docker-compose.yml`内の`command: go run main.go`および`command: npm run dev`のコメントアウトを解除します。

以下のコマンドを実行し、コンテナを起動します。

```
$ docker-compose up -d
```

## コンテナ停止

以下のコマンドを実行します。

```
$ docker-compose down
```
