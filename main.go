package main

import (
	"embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	app := NewApp()

	// コマンドライン引数からファイルパスを取得（ダブルクリック起動対応）
	if len(os.Args) > 1 {
		filePath := os.Args[1]
		ext := strings.ToLower(filepath.Ext(filePath))
		if ext == ".md" || ext == ".markdown" || ext == ".txt" {
			app.initialFilePath = filePath
		}
	}

	err := wails.Run(&options.App{
		Title:  "Markdown Viewer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 27, B: 27, A: 0},
		OnStartup:        app.startup,
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop:     true,
			DisableWebViewDrop: true, // ★修正: WebViewのデフォルトドロップ動作を無効化
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
