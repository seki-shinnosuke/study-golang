# パッケージ/ディレクトリ構成
```
src
├── cmd             //各アプリケーションのmain関数を配置
│   ├── api
│   └── batch
├── code            //ENUMチックなコード定義を配置
├── config
│   ├── config.go   //Viperを利用した環境変数の設定
│   └── database.go //RDBの接続設定
├── controller      //各機能のコントローラーを配置
├── error           //カスタマイズしたエラーコード
├── model
│   ├── db          //SQLBoilerを利用したORM
│   ├── graphql     //TODO
│   └── rest        //REST APIのリクエストレスポンス構造を配置
├── server
│   ├── server.go   //Ginを利用したREST APIの設定
│   └── wire.go     //wireを利用したDIの設定
├── usecase         //各機能のビジネスロジックを配置
└── util            //共通ロジックを配置
```

# REST API開発の流れ
1. APIのビジネスロジック/コントローラーを作成  
2. server.goのNewServerに作成したコントローラーの構造体を設定  
3. server.goのsetRoutingにAPIのルーティングを設定  
4. wire.gonのInitializeServiceにコントローラーの構造体を設定し`make wire-api`でDI関数を自動生成

# GraphQL 開発の流れ
TODO