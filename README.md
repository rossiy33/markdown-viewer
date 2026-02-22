# Markdown Viewer

Go + Wails v2 で作成したシンプルなMarkdownビューアです。

![Windows](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-blue)
![Go](https://img.shields.io/badge/Go-1.25-00ADD8)
![Wails](https://img.shields.io/badge/Wails-v2.11-red)

## 特徴

- GitHub風のMarkdownレンダリング（GFM対応）
- ダークモード / ライトモード自動切り替え（OS設定に追従）
- ドラッグ＆ドロップでファイルを開く
- `.md` `.markdown` `.txt` ファイルに対応
- シンタックスハイライト（highlight.js）
- 単一バイナリで動作（外部依存なし）

## ダウンロード

[Releases](https://github.com/rossiy33/markdown-viewer/releases) から各OS向けのバイナリをダウンロードできます。

| OS | ファイル |
|---|---|
| Windows | `markdown-viewer-windows-amd64.exe` |
| macOS (Universal) | `markdown-viewer-darwin-universal.zip` |
| Linux | `markdown-viewer-linux-amd64` |

## 使い方

1. アプリを起動する
2. 以下のいずれかの方法でファイルを開く：
   - **「ファイルを開く」ボタン**をクリック
   - ファイルをウィンドウに**ドラッグ＆ドロップ**
   - Welcome画面を**ダブルクリック**
   - `.md` ファイルの関連付けで**直接起動**

## ビルド

### 前提条件

- Go 1.21以上
- [Wails CLI v2](https://wails.io/docs/gettingstarted/installation)

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### ビルド手順

```bash
git clone https://github.com/rossiy33/markdown-viewer.git
cd markdown-viewer
wails build
```

実行ファイルは `build/bin/` に生成されます。

### 開発モード

```bash
wails dev
```

## ライセンス

MIT
