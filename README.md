# GraghQLserver-with-Golang
Go言語を利用してGraghQLについて学ぶ

# GraghQLとは ？
* RESTful Web APIの持つ問題を解決するために開発された規格
* Facebookが開発しているWebAPIのための規格
* 「クエリ言語」と「スキーマ言語」からなる
* スキーマファースト開発ができる
* マイクロサービスとも相性が良い
* GraphQLはマイクロサービスのGatewayとしても良い役割を発揮する

## クエリ言語とは？
* GraphQL APIのリクエストのための言語
* データ取得系のQuery, データ更新系のQuery, イベントの通知であるSubscriptionがある

## スキーマ言語とは?
* GraphQL APIの仕様を記述するための言語
* リクエストされたクエリはスキーマ言語で記述したスキーマに従ってGraphQL処理系によって実行されレスポンスを生成する

## スキーマファースト開発とは？
* API開発者とフロントエンド開発者の間で事前にAPIのスキーマを相談して決定し、パラレルに開発することで結合時の手戻りを防ぐプロセス

## 特徴
* クエリがレスポンスデータの構造と似ていて情報量が多い
* スキーマによる型付けにより型安全な運用ができる
* 必要なデータだけリクエストして受け取る
* RESTのようなN回/時間で制限をかけることができない
* エラーハンドリングが難しい、RESTではHTTPステータスコードがエラーを表現してくれる
* 単一のエンドポイントでエンドポイントごとにレスポンス性能を監視することができない
* 必要な情報しか通信が行われないため余分な通信が行われない
* フロンドエンド側で取得する情報が決定できるため、開発が早くなる
* スキーマによる型定義の恩恵を得られる
 * スキーマによってどのようなフィールドがあるのか、どのような形で取得できるのかが明確になる

## REST(REpresentational State Transfer)とは
* WebAPIの設計モデル
* WebAPIの仕様を決める上でのアーキテクチャスタイル
* GraphQLは言語, RESTはアーキテクチャ
* サービスのURIにHTTPメソッドでアクセスすることでデータの送受信を行う

# Go言語でGraphQLサーバーを作成
## gqlgenとは？
* 現在Go言語でGraphQLサーバーを実装する手段は下記2点
 * graphql-go/graphql (Star 6.2k)
 * 99designs/gqlgen (Star 4.2k)

## Go言語について
* 最新版へのアップデート
 * brew upgrade go

## パッケージ管理ツール
* goコマンドにも組み込まれている```go mod```が使用可能
 * ```go mod init```で初期化する
 * ```go build```などのビルドコマンドで依存モジュールを自動インストール
 * ```go list -m all```で現在の依存モジュールを表示する
 * ```go get```で依存モジュールの追加やバージョンアップを行う
 * ```go mod tidy```で使われていない依存モジュールを削除する
* Go Modulesを使うのがデファクト
 * https://github.com/golang/dep (Star：13.2k)
 * https://github.com/golang/gomodule (Go言語のメインリポジトリに統合された)

 ## 実装
 * パブリックIPでCloudSQLを立ち上げ  
 → 承認済みネットワークに自宅IPを設定
 * データベース作成  
   ```gcloud sql databases create hayashi-golang --instance=hayashi-golang```
 * データベース接続  
   ```gcloud sql connect hayashi-golang --user=postgres --quiet```
   → db/schema.dbのテーブルを作成    
 * データベース接続関数：connectDatabase()
 * GraphQLサーバーとするための初期化を実行
    ```gqlgen init```
    → ```server.go```と```gragh```が作成される
 * GraphQLスキーマ(graph/schema.graphqls)はサンプルであるため、修正してソースコードを再生成
    ```rm graph/schema.resolvers.go``` → ```gqlgen```

# 歴史
## Webの黎明期
* 初期のWebサイトは性的なHTMLドキュメントがインターネット越しに送られてくるだけのシンプルなものだった
* SQLなどでデータベースに格納されたコンテンツを動的に取得したり、Javascriptで操作するようになった

## REST:APIの隆盛
* 開発者たちはすべての形、サイズのアプリでデータを表示するためにRESTfulAPIを利用し始めた

## GraphQL：APIの進化系
* API作成の仕組みとしてRESTの代わりに使用できる
* GraphQLは標準化された言語、型付け、仕様を持ちクライアントとサーバー間を協力に結びつける
* 異なるデバイス間の通信に標準化された言語であり、大型かつクロスぷらっとフォームのアプリ開発がよりシンプルになる

# 参考
* https://qiita.com/SiragumoHuin/items/cc58f456bc43a1be41b4
* https://rightcode.co.jp/blog/information-technology/graphql-alternative-rest-api
* https://employment.en-japan.com/engineerhub/entry/2018/12/26/
* https://tech.opst.co.jp/2019/07/09/go-modules%E3%82%82%E8%A7%A6%E3%82%8C%E3%81%A6%E3%81%BF%E3%82%8Bgo%E5%85%A5%E9%96%80/
 * GOPATHやGO MODULESについて整理されてる