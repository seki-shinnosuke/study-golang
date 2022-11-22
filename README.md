# study-golang
📚Go言語でのWebアプリケーション学習用プロジェクト📚

本プロジェクトでは以下の技術要素を利用しGolangによるWebアプリケーション開発の基礎を学んで頂くことを目的とします

| 技術 | バージョン | 用途 |
| :---: | :---: | :---: |
| Golang | 1.19 | サーバ言語 |
| Gin | 1.8.1 | サーバ言語FW |
| MySQL | 8.0.31 | DB |
| Air | 1.40.4 | ホットリロード |
| Docker Compose| >=2.10.2 | コンテナ(DB構築用) |

## 事前準備
開発環境では以下アプリケーションをインストールしてください
1. Visual Studio Code & Plugin<br />
https://code.visualstudio.com/<br />
https://marketplace.visualstudio.com/items?itemName=golang.Go
2. Docker Desktop<br />
https://www.docker.com/products/docker-desktop

本プロジェクトはDocker Composeを利用しGolangおよびミドルウェアの依存関係を全てコンテナに集約する方式を採用しています  
```
% make build
```

## 実行
以下のコマンドでアプリケーションを実行します  
ローカルで変更されたコードはホットリロードによりコンテナ内のアプリケーションに反映されます  
```
% make up
```

## 停止
```
% make down
```
その他のコマンドは[Makefile](./Makefile)を参照してください
