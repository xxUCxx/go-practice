# Go Todo API (`go-practice`)

Go言語の標準ライブラリ（`net/http`, `encoding/json`）のみを使用して構築した、タスク管理（Todo）用のWeb APIです。  
外部フレームワークに頼らず、HTTPリクエストのルーティング、JSONのシリアライズ／デシリアライズ、ステータスコード制御の基礎を実装しています。

## 使用技術
- **Language**: Go 1.23
- **Infrastructure**: Docker / Docker Compose
- **Version Control**: Git / GitHub（ブランチ運用・Pull Requestベースの開発）

## ファイル構成
```text
.
├── compose.yml   # Go実行環境のDocker Compose設定
├── go.mod        # Goモジュール定義
├── main.go       # Webサーバー・Todo APIの実装本体
└── README.md     # 本ドキュメント
```

## 起動・停止方法

Docker環境があれば、以下のコマンドのみでサーバーを起動できます。

```bash
# バックグラウンドでサーバーを起動
docker compose up -d

# サーバーを停止
docker compose down
```

## API エンドポイント仕様

### 1. Todo一覧の取得
登録されているすべてのTodoをJSON配列で返します。

- **Method**: `GET`
- **Endpoint**: `/todo`
- **リクエスト例**:
  ```bash
  curl http://localhost:8080/todo
  ```
- **レスポンス例 (`200 OK`)**:
  ```json
  [
    {
      "id": 1,
      "title": "Goを復習する",
      "is_done": false
    }
  ]
  ```

### 2. Todoの新規登録
JSONデータを受け取り、IDを自動採番してリストに追加します。

- **Method**: `POST`
- **Endpoint**: `/todo`
- **リクエスト例**:
  ```bash
  curl -X POST http://localhost:8080/todo \
    -H "Content-Type: application/json" \
    -d '{"title":"POSTの復習をする","is_done":false}'
  ```
- **レスポンス例 (`201 Created`)**:
  ```json
  {
    "id": 2,
    "title": "POSTの復習をする",
    "is_done": false
  }
  ```

## 今後の実装予定（Roadmap）
- [x] Todo一覧取得（`GET`）・新規登録（`POST`）の実装
- [ ] ハンドラ関数の切り出しによるリファクタリング
- [ ] 指定IDのTodo削除（`DELETE`）・更新（`PUT`）の実装
- [ ] データベース（MySQL）連携によるデータの永続化
