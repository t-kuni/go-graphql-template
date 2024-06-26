# Go GraphQL Template

This repository is project template for Go GraphQL application.

# Features

* [Onion Architecture](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/)
* [DI Container](https://github.com/samber/do)
* Server generation from [GraphQL Scheme](https://graphql.org/learn/schema/) with [gqlgen](https://gqlgen.com/)
* [Validator](https://github.com/go-playground/validator)
* [ORM](https://github.com/ent/ent)
* [Logging](https://github.com/sirupsen/logrus)
* [Error Handling (Stack trace)](https://github.com/rotisserie/eris)
* Seeder

# Requirements

* go 1.19+
  * [How to install and switch between multiple versions of golang](https://gist.github.com/t-kuni/4e23b59f16557d704974b1ce6b49e6bb)

# Usage

```
cp .env.example .env
cp .env.feature.example .env.feature
make generate
docker compose up -d
```

DB Migration and Seeding

```
docker compose exec app sh
go run commands/migrate/main.go
go run commands/seed/main.go
```

Confirm

```
http://localhost:34567/
```

# Tests

Unit test

```
go test ./...
```

Feature test  
https://localhost:8080 に接続し、`example_test`データベースを作成してから以下のコマンドを実行する

```
docker compose exec app sh
DB_DATABASE=example_test go run commands/migrate/main.go
gotestsum --hide-summary=skipped -- -tags feature ./...
```

# Setting remote debug on GoLand

https://gist.github.com/t-kuni/1ecec9d185aac837457ad9e583af53fb#golnad%E3%81%AE%E8%A8%AD%E5%AE%9A

# See Database

http://localhost:8080

# See SQL Log

```
docker compose exec db tail -f /tmp/query.log
```

# Create Scheme

```
go run entgo.io/ent/cmd/ent init [EntityName]
```

# Build Container for production

```
docker build --target prod --tag go-graphql-template .
```

# タスク

- [ ] 自動発行されるIDをUUIDにする
- [ ] graphql-playgroundを別コンテナに切り出す
- [ ] クエリの絞り込み処理
- [x] 認証処理
- [ ] 認可処理
  - [x] Directiveを実装してみる
  - [ ] contextを見ればどのqueryが呼ばれたかなど分かるのか？
- [ ] Subscribeを試す
- [ ] polluterからtestfixturesに載せ替え
- [ ] テストを書けるようにする
- [ ] マイグレーションの管理を切り出し
- [ ] 認証処理のモック化
- [ ] レスポンスがJSONではない処理のテスト（例えばファイルのダウンロードなど）
- [ ] 現在日時のモック化
- [ ] DB接続のタイムゾーン
- [ ] 本番環境用コンテナ
- [ ] vscode用devcontainer定義
- [ ] coreファイルが残る問題
- [ ] アクセスログ
- [ ] テストコードの用意
- [ ] CI
- [ ] CD
