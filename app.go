package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx             context.Context
	initialFilePath string // ダブルクリック起動時のファイルパス
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// ファイルドロップイベントのハンドラを登録
	runtime.OnFileDrop(a.ctx, func(x, y int, paths []string) {
		if len(paths) == 0 {
			return
		}

		for _, path := range paths {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".md" || ext == ".markdown" || ext == ".txt" {
				content, err := os.ReadFile(path)
				if err != nil {
					runtime.EventsEmit(a.ctx, "file-error", err.Error())
					return
				}
				runtime.EventsEmit(a.ctx, "file-loaded", map[string]string{
					"path":    path,
					"content": string(content),
				})
				return
			}
		}
		runtime.EventsEmit(a.ctx, "file-error", "Markdownファイル (.md, .markdown, .txt) を選択してください")
	})
}

// GetInitialFile はコマンドライン引数で渡されたファイルを返す（ダブルクリック起動対応）
func (a *App) GetInitialFile() (map[string]string, error) {
	if a.initialFilePath == "" {
		return nil, nil
	}

	content, err := os.ReadFile(a.initialFilePath)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"path":    a.initialFilePath,
		"content": string(content),
	}, nil
}

// ReadMarkdownFile はパスを受け取ってMarkdownファイルを読み込む（JS側のOnFileDropから直接呼ばれる）
func (a *App) ReadMarkdownFile(path string) (map[string]string, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".md" && ext != ".markdown" && ext != ".txt" {
		return nil, fmt.Errorf("Markdownファイル (.md, .markdown, .txt) を選択してください")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"path":    path,
		"content": string(content),
	}, nil
}

// OpenFileDialog はファイル選択ダイアログを開いてMarkdownファイルを読み込む
func (a *App) OpenFileDialog() (map[string]string, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Markdownファイルを開く",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files (*.md, *.markdown)",
				Pattern:     "*.md;*.markdown",
			},
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"path":    path,
		"content": string(content),
	}, nil
}
