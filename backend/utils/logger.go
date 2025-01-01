package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ログの設定
func SetupLogger() error {

	logDir := "logs"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("ログディレクトリの作成に失敗しました: %v", err)
		}
	}

	logFile, err := os.OpenFile(logDir+"/application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("ログファイルのオープンに失敗しました: %v", err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return nil
}
