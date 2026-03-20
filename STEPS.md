# kohanDB - 実装ステップ

## Phase 1: CREATE TABLE を動かす

- [x] Step 1 — プロジェクト初期化 (`go mod init`, `main.go`)
- [x] Step 2 — REPL (対話ループ)
- [x] Step 3 — Tokenizer (字句解析)
- [ ] Step 4 — Parser (構文解析)
- [ ] Step 5 — Catalog (テーブル定義のメモリ管理)
- [ ] Step 6 — ファイルへの永続化

## Phase 2以降 (ロードマップ)

- [ ] Phase 2 — INSERT (行データのメモリ格納 + ファイル永続化)
- [ ] Phase 3 — SELECT * (全行スキャン + 結果表示)
- [ ] Phase 4 — WHERE句 (式の評価・比較演算)
- [ ] Phase 5 — DELETE / UPDATE
- [ ] Phase 6 — ページベースのストレージ (B-Tree導入)
- [ ] Phase 7 — より本格的なSQL対応 (JOIN, ORDER BY, etc.)
