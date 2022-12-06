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
### 言語
- Go
- Gin
- GORM

### データベース
- MySQL?

### インフラ
- AWS Lamda?
- API Gateway?
- Docker
- docker-compose

### 外部API
- LINE Messaging API
- 天気API（何を使うかは未定）

## 4.テーブル設計

## 5.エンドポイント
| 目的                       | メソッド | エンドポイント   | 
| -------------------------- | -------- | ---------------- | 
| Trashdayの一覧を取得       | GET      | /trashdays       | 
| Trashdayのデータを登録     | POST     | /trashdays       | 
| Trashdayのデータ１件を取得 | GET      | /trashdays/[:id] | 
| Trashdayデータを更新       | PUT      | /trashdays/[:id] | 
| Trashdayデータを削除       | DELETE   | /trashdays       | 

## 6.起動方法


## 7.参考
- [Go(Golang)+Herokuで天気予報通知LINE BOTを作成する](https://qiita.com/yuki_0920/items/cbdbd5220a6a8b4eef19)
- [AWS Lambdaを利用したLINEbotハンズオン](https://cloud5.jp/aws-lambda_line-api/)
- [【初心者向け】GASでゴミ出しLINE Botをつくるための「チャネル」とその作成](https://tonari-it.com/gas-line-bot-create-channel/)
- [人気の天気APIをまとめてみた](https://qiita.com/cnakano/items/ff3fd90f685f4ca363cc)