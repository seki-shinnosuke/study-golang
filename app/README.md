# パッケージ/ディレクトリ構成
```
src
├── cmd           //各アプリケーションのmain関数を配置
│   ├── api
│   └── batch
├── config
│   └── config.go //Viperを利用した環境変数の設定
├── model         //TODO
├── server
│   ├── <service> //各機能のビジネスロジックを配置
│   ├── server.go //Ginを利用したREST APIの設定
│   └── wire.go   //wireを利用したDIの設定
└── util          //共通ロジックを配置
```

# REST API開発の流れ
1. APIのビジネスロジックを作成  
2. server.goのNewServerに作成したロジックの構造体を設定  
3. server.goのsetRoutingにAPIのルーティングを設定  
4. wire.gonのInitializeServiceにロジックの構造体を設定し`make wire`でDI関数を自動生成
  