package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// LZHファイルのダウンロードと展開
func DownloadAndExtractLzh(url string, extractPath string) error {
	// 一時ファイル名
	tempFile := "downloaded.lzh"

	// ファイルのダウンロード
	err := downloadFile(url, tempFile)
	if err != nil {
		return fmt.Errorf("ダウンロードファイルが見つかりません: %v", err)
	}

	// 7zを使ってLZHファイルを解凍
	err = extractLzhWith7z(tempFile, extractPath)
	if err != nil {
		return fmt.Errorf("7zを使ったLZHファイルの解凍に失敗しました: %v", err)
	}

	// ダウンロードした一時ファイルの削除
	err = os.Remove(tempFile)
	if err != nil {
		log.Println("ダウンロードしたファイルの削除に失敗しました:", err)
	}

	return nil
}

// ファイルをダウンロードする
func downloadFile(url string, filepath string) error {
	// URLからHTTP GETリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ダウンロードに失敗しました: %v", err)
	}
	defer resp.Body.Close()

	// ダウンロードしたファイルを保存する
	outFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("ファイルの作成に失敗しました: %v", err)
	}
	defer outFile.Close()

	// ファイル内容をコピー
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("ファイルのコピーに失敗しました: %v", err)
	}

	return nil
}

// 7zを使ってLZHファイルを解凍する
func extractLzhWith7z(lzhFile string, extractPath string) error {
	// 解凍先ディレクトリの作成
	err := os.MkdirAll(extractPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ディレクトリの作成に失敗しました: %v", err)
	}

	// 7zコマンドを使用してLZHファイルを解凍
	cmd := exec.Command("7z", "x", lzhFile, "-o"+extractPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("7zコマンドの実行に失敗しました: %v", err)
	}

	return nil
}
