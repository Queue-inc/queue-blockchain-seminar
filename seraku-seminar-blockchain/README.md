# Bitcoin-vive

## How to use

## api

### ユーザー登録
- 100COINもらえます。
```json
{
  "body": {
    "type": "create_user",
    "entity": {
      "name": "ユーザー名",
      "_id": "bson.ObjectId"
    }
  },
  "signature": "電子署名",
  "public_key": "公開鍵"
}
```

### 送金
```json
{
  "body": {
    "type": "send_funds",
    "entity": {
      "to": "送り先の公開鍵",
      "amount": "送りたい量(Integer)"
    }
  },
  "signature": "電子署名",
  "public_key": "公開鍵"
}
```