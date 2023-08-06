# 使い方
## 前提
- Goのバージョン1.16以上がインストールされている必要があります
- npmがインストールされている必要があります

## 必要なパッケージのインストール
#### バックエンド
```bash
cd grpc-back
make go-tools-install
make set-path
```
#### フロントエンド
```bash
cd grpc-front
make node-tools-install
make set-path
```

## Protocol Buffersの生成
```bash
cd proto
make proto-gen
```

## Modelsの生成
- Modelsに対応するスキーマがマイグレーション済みである必要があります
```bash
cd grpc-back
make models-gen
```

## アプリの立ち上げ
#### バックエンド
```bash
cd grpc-back
docker compose up -d
```
#### フロントエンド
```bash
cd grpc-front
docker compose up -d
```

## マイグレーション
- db-migrationsコンテナを再起動してください
- コンテナ経由ではなく、手動で行う場合は以下の通りです
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
cd grpc-back/db

# マイグレーションアップ
make up

# マイグレーションダウン
make down
```