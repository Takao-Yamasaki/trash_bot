# ゴミ出しLINE Botの作成

## 1.開発の背景
- 私自身ゴミの回収日を把握していないため、ゴミの出し忘れをすることがよくある。
- ゴミ出しの準備をしていても、当日雨が降っていて、ゴミ出しできないこともある。
- 「今日の天気」と「今日の収集物」について、自分自身よく使うLINEに通知することでこれらの課題感を解決したいと考えた。

## 2.実装する機能
- 今回、実装する機能については、以下のとおり。
- タスク管理については、[Notion](https://separate-decade-f0a.notion.site/382acc6839c348c9a10b9041d163b33c?v=1352c1b85d8544789e0ffad9c0e6524d) を使用。

### ユーザー側
- 朝8時に「今日の天気」「今日の収集物」「今日のコメント」をLINEに通知する
    - 「今日の天気」のLINE通知
    - 「今日の収集物」のLINE通知

### 管理者側
- ゴミ出しの内容（収集日と収集物）について、Web上の管理画面から登録、取得、削除、更新、削除処理ができる
    - ゴミ出しの内容（収集日と収集物）のCRUD処理
    - 管理者ログイン機能（今後実装）

## 3.使用技術
### 言語・フレームワーク
- Go(1.19)
- Gin(1.8.1)
- Gorm(1.24.2)

### データベース
- SQLite3(3.37.0)

### インフラ
- Docker(未実装)
- docker-compose(未実装)

### 外部API
- LINE Messaging API(未実装)
- Open Weather(未実装)

## 4.構成図
![infra](https://user-images.githubusercontent.com/24619682/206127178-5a89428a-1cd3-48e0-8fb3-79ea2feaf273.png)

## 5.テーブル設計
![table](https://user-images.githubusercontent.com/24619682/206127671-2926b827-69eb-4208-badb-3bb74b4810c3.png)

## 5.エンドポイント
| 目的                       | メソッド | エンドポイント    | 
| -------------------------- | -------- | ----------------- | 
| TrashDayの一覧を取得       | GET      | /trash-days       | 
| TrashDayのデータ１件を取得 | GET      | /trash-day/:id    | 
| TrashDayのデータを登録     | POST     | /trash-day/create | 
| TrashDayデータを更新       | POST     | /trash-day/update | 
| TrashDayデータを削除       | POST     | /trash-day/delete | 
| Adminの一覧を取得          | GET      | /admins           | 
| Adminのデータ１件を取得    | GET      | /admin/:id        | 
| Adminのデータを登録        | POST     | /admin/create     | 
| Adminデータを更新          | POST     | /admin/update     | 
| Adminデータを削除          | POST     | /admin/delete     | 
| Commentの一覧を取得        | GET      | /comments         | 
| Commentのデータ１件を取得  | GET      | /comment/:id      | 
| Commentのデータを登録      | POST     | /comment/create   | 
| Commentデータを更新        | POST     | /comment/update   | 
| Commentデータを削除        | POST     | /comment/delete   | 
- 更新処理と削除処理のHTTPメソッドについては、今後見直す予定

## 6.起動方法
- Dockerfile/docker-compose.ymlについては、今後実装予定
```
$ docker-compose up -d
```

## 7.参考
- [Go(Golang)+Herokuで天気予報通知LINE BOTを作成する](https://qiita.com/yuki_0920/items/cbdbd5220a6a8b4eef19)
- [AWS Lambdaを利用したLINEbotハンズオン](https://cloud5.jp/aws-lambda_line-api/)
- [【初心者向け】GASでゴミ出しLINE Botをつくるための「チャネル」とその作成](https://tonari-it.com/gas-line-bot-create-channel/)
- [人気の天気APIをまとめてみた](https://qiita.com/cnakano/items/ff3fd90f685f4ca363cc)
- [GoとLambdaでAWS使用料金を毎日LINE通知させる](https://qiita.com/yu9penguin/items/8657069a5d270d3c4c9c)