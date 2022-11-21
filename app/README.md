# パッケージ/ディレクトリ構成
```
src
├── cmd             //各アプリケーションのmain関数を配置
│   ├── api
│   └── batch
├── config
│   ├── config.go   //Viperを利用した環境変数の設定
│   └── database.go //RDBの接続設定
├── model         
│   └── db          //SQLBoilerを利用したORM
├── controller      //各機能のコントローラーを配置
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
4. wire.gonのInitializeServiceにコントローラーの構造体を設定し`make wire`でDI関数を自動生成
  