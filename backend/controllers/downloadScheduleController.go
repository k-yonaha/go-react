package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"backend/database"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func DownloadSchedule(c echo.Context) error {

	if err := utils.SetupLogger(); err != nil {
		log.Fatalf("ログの設定に失敗しました: %v", err)
	}
	// クエリパラメータから日付を取得
	date := c.QueryParam("date")
	if date == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "日付パラメータが必要です")
	}

	// URLを構築（例: https://www1.mbrace.or.jp/od2/B/202412/b241225.lzh）
	// date[:6] 最初の６文字取得
	// date[2:] 最初の２文字を削除して残りを取得
	url := fmt.Sprintf("https://www1.mbrace.or.jp/od2/B/%s/b%s.lzh", date[:6], date[2:])

	// 解凍先のパス
	extractPath := fmt.Sprintf("./extracted_files/%s", date)

	// 解凍先ディレクトリが既に存在する場合は、解凍処理をスキップ
	if _, err := os.Stat(extractPath); !os.IsNotExist(err) {
		log.Printf("ファイルは既に解凍されています: %s", extractPath)
	} else {
		// LZHファイルのダウンロードと解凍処理
		err := utils.DownloadAndExtractLzh(url, extractPath)
		if err != nil {
			log.Println("エラー:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "LZHファイルのダウンロードまたは解凍に失敗しました")
		}
	}

	// 入力された日付のデータが存在する場合は処理しない
	convertedDate := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
	parsedDate, err := time.Parse("2006-01-02", convertedDate)
	if err != nil {
		log.Printf("日付の解析に失敗しました: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "日付の形式が正しくありません")
	}
	utcDate := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.UTC)
	exists, err := services.RaceScheduleExists(database.DB, utcDate)
	if err != nil {
		return fmt.Errorf("レーススケジュールの確認に失敗しました: %v", err)
	}

	if exists {
		log.Printf("レース日 %s のスケジュールはすでに存在します。処理をスキップします。", utcDate.Format("2006-01-02"))
		// 重複している場合は処理しない
	} else {

		// 解凍したファイルのパスを取得
		files, err := os.ReadDir(extractPath)
		if err != nil {
			log.Println("ファイル読み込みエラー:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "解凍後のファイルの読み込みに失敗しました")
		}

		// 解凍したファイルからレーススケジュールを読み込んでデータベースに保存
		for _, file := range files {
			if !file.IsDir() {
				filePath := filepath.Join(extractPath, file.Name())

				// 解凍されたファイルの内容を読み込む（例: .txt ファイル）
				err := processFile(c, filePath)
				if err != nil {
					log.Println("ファイル読み込みエラー:", err)
					return echo.NewHTTPError(http.StatusInternalServerError, "レース情報の保存に失敗しました")
				}

			}
		}
	}
	return c.String(http.StatusOK, fmt.Sprintf("LZHファイル %s のダウンロード、解凍、解析、保存が成功しました", date))
}

// レース情報の解析とデータベース保存を行う
func DownloadAndSaveRaceSchedule(c echo.Context, rawData []byte) error {
	// レースのセクションを分割
	rawDataStr := string(rawData)

	re := regexp.MustCompile(`\d+BBGN`)
	sections := re.Split(rawDataStr, -1)
	var raceSchedules []models.RaceSchedule
	for _, section := range sections {
		// 各レース情報を解析
		if len(section) > 0 {
			parsedSchedules, err := parseRaceSchedule(section)

			if err != nil {
				return fmt.Errorf("レース情報の解析に失敗しました。: %v", err)
			}
			raceSchedules = append(raceSchedules, parsedSchedules...)
		}
	}

	// データベースに保存

	for _, raceSchedule := range raceSchedules {
		err := services.CreateRaceSchedule(database.DB, raceSchedule)
		if err != nil {
			return fmt.Errorf("レース情報の保存に失敗しました。: %v", err)
		}
	}

	return nil
}

// レース情報のパース関数
func parseRaceSchedule(data string) ([]models.RaceSchedule, error) {
	var raceSchedules []models.RaceSchedule
	raceSchedule := models.RaceSchedule{}
	var raceScheduleParticipants []models.Participant
	lines := strings.Split(data, "\n")
	// レースの日付と時間、レース番号を抽出
	for i, line := range lines {

		if strings.Contains(line, "STARTB") || line == "" {
			continue
		}

		if strings.Contains(line, "END") || strings.Contains(line, "FINALB") {
			continue
		}

		// ボートレース場を抽出
		if i == 1 {
			raceSchedule.RaceDay = extractRaceDay(line)
			raceSchedule.CourseName = extractCourseName(line)
		}

		// 番組名
		if i == 3 {
			raceSchedule.RaceProgram = extractRaceProgram(line)

		}

		// レース日時
		if i == 4 {
			raceSchedule.RaceDate = extractRaceDate(line)
		}

		// 締切時刻
		if i == 6 || strings.Contains(line, "電話投票締切予定") {

			// 1R,2R
			raceSchedule.RaceNumber = extractRaceNumber(line)
			// 予選,特選
			raceSchedule.RaceType = extractRaceType(line)
			// 投票締切時刻
			raceSchedule.RaceTime = extractRaceTime(line, raceSchedule.RaceDate)

		}

		// 選手情報を取得（例: 選手番号、名前、級）
		if len(line) > 0 && line[0] >= '1' && line[0] <= '6' {
			// 選手情報を処理
			participants, err := extractParticipants(line)
			if err != nil {
				return nil, fmt.Errorf("選手情報の抽出失敗: %v", err)
			}
			raceScheduleParticipants = append(raceScheduleParticipants, participants...)

			if line[0] == '6' {
				participantsJson, err := json.Marshal(raceScheduleParticipants)
				if err != nil {
					log.Printf("Participants を JSON に変換できませんでした: %v", err)
					return nil, fmt.Errorf("選手情報の抽出失敗: %v", err)
				}

				raceSchedule.Participants = participantsJson
				raceSchedules = append(raceSchedules, raceSchedule)

				// raceSchedule = models.RaceSchedule{}
				raceScheduleParticipants = nil
			}
		}

	}

	return raceSchedules, nil
}

// 「第 3日」などの日程を抽出
func extractRaceDay(line string) string {
	line = strings.ReplaceAll(line, "　", "")
	// 正規表現を使用して「第 X日」を抽出
	parts := strings.Fields(line)

	if len(parts) > 0 {
		return parts[3]
	}
	return ""
}

// ボートレース場を抽出
func extractCourseName(line string) string {
	// line = toHalfWidth(line)
	line = strings.ReplaceAll(line, "　", "")
	parts := strings.Fields(line)
	if len(parts) > 0 {
		courseName := parts[0]
		courseName = strings.Replace(courseName, "ボートレース", "", -1)
		courseName = strings.Replace(courseName, " ", "", -1)
		return courseName
	}
	return ""
}

func extractRaceProgram(line string) string {
	parts := strings.Fields(line)
	if len(parts) > 0 {
		raceProgram := parts[0]
		return raceProgram
	}
	return ""
}

// レース日付を抽出
func extractRaceDate(line string) time.Time {
	// "２０２４年１２月２３日" のような形式を処理
	dateParts := strings.Fields(line)
	if len(dateParts) >= 3 {
		// 全角を半角にする
		raceDate := toHalfWidth(dateParts[2])

		raceDate = strings.ReplaceAll(raceDate, "年", "-")
		raceDate = strings.ReplaceAll(raceDate, "月", "-")
		raceDate = strings.ReplaceAll(raceDate, "日", "")

		parsedDate, err := time.Parse("2006-01-02", raceDate)

		if err != nil {
			log.Printf("日付の解析に失敗しました: %v", err)
			return time.Time{} // エラー時はゼロ値の time.Time を返す
		}
		return parsedDate
	}
	return time.Time{} // 日付が見つからない場合はゼロ値の time.Time を返す
}

// レース時間を抽出
func extractRaceTime(line string, raceDate time.Time) time.Time {
	line = strings.ReplaceAll(line, "　", "")
	line = strings.ReplaceAll(line, "進入固定", "")
	timeParts := strings.Fields(line)

	if len(timeParts) > 1 {
		rawTime := timeParts[3]
		// 文字列を半角に変換
		halfWidthStr := toHalfWidth(rawTime)
		raceTime, err := extractTime(halfWidthStr)

		if err != nil {
			log.Printf("時間の抽出に失敗しました: %v", err)
			return time.Time{}
		}

		parsedTime, err := time.Parse("15:04", raceTime)
		if err != nil {
			log.Printf("時間のフォーマットに失敗しました1: %v", err)
			return time.Time{}
		}

		dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d", raceDate.Year(), raceDate.Month(), raceDate.Day(), parsedTime.Hour(), parsedTime.Minute())

		parsedDate, err := time.Parse("2006-01-02 15:04", dateStr)
		if err != nil {
			log.Printf("時間のフォーマットに失敗しました2: %v", err)
			return time.Time{}
		}

		return parsedDate
	}
	return time.Time{}
}

// 予選、特選を抽出
func extractRaceType(line string) string {
	line = strings.ReplaceAll(line, "　", "")
	parts := strings.Fields(line)

	if len(parts) > 1 {
		// 文字列を半角に変換
		halfWidthStr := toHalfWidth(parts[1])
		if parts[2] == "進入固定" {
			halfWidthStr = halfWidthStr + " " + parts[2]
		}

		return halfWidthStr
	}
	return "" // 時間が見つからない場合はゼロ値の time.Time を返す
}

func extractTime(str string) (string, error) {
	// "電話投票締切予定" の後の時間を取得
	re := regexp.MustCompile(`電話投票締切予定(\d{2}:\d{2})`)
	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("時間が見つかりませんでした")
}

// レース番号を抽出
func extractRaceNumber(line string) int {
	parts := strings.Fields(line)
	if len(parts) > 0 {
		// 例: "１Ｒ" から "1" を抽出
		halfWidthStr := toHalfWidth(parts[0])
		halfWidthStr = strings.ReplaceAll(halfWidthStr, "R", "")

		if num, err := strconv.Atoi(halfWidthStr); err == nil {
			return num
		}
	}
	return 0
}

// 選手情報を抽出
func extractParticipants(line string) ([]models.Participant, error) {

	re := regexp.MustCompile(`^(\d+)\s+(\d+)([^\d]+)(\d{2})([^\d]+)(\d+[A-Za-z0-9]+)`)
	matches := re.FindStringSubmatch(line)

	// マッチする部分を処理
	if len(matches) < 7 {
		return nil, fmt.Errorf("選手情報の抽出失敗 不正なフォーマット: %s", line)
	}

	age, err := strconv.Atoi(matches[4])
	if err != nil {
		log.Printf("年齢の解析に失敗しました: %v", err)
		age = 0 // エラー時は0歳
	}
	// 選手情報の形式に合わせて処理する
	// 選手情報の形式に合わせて処理する
	participant := models.Participant{
		CourseNumber: matches[1], // 艇番
		PlayerNumber: matches[2], // 選手番号
		Name:         matches[3], // 名前
		Age:          age,        // 年齢
		Region:       matches[5], // 支部
		Grade:        matches[6], // 体級種別
	}

	return []models.Participant{participant}, nil
}

func convertShiftJISToUTF8(filePath string) ([]byte, error) {
	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開くことができません: %v", err)
	}
	defer file.Close()

	// Shift_JIS エンコーディングを UTF-8 に変換
	decoder := japanese.ShiftJIS.NewDecoder()
	utf8Reader := transform.NewReader(file, decoder)

	// 変換されたデータを読み込み
	convertedData, err := io.ReadAll(utf8Reader)
	if err != nil {
		return nil, fmt.Errorf("Shift_JIS から UTF-8 に変換中にエラーが発生しました: %v", err)
	}

	return convertedData, nil
}

func processFile(c echo.Context, filePath string) error {
	// 文字化けしたファイルを読み込んで変換
	convertedData, err := convertShiftJISToUTF8(filePath)
	if err != nil {
		log.Fatalf("ファイルの変換エラー: %v", err)
	}

	// 変換後のデータを行単位で処理
	lines := strings.Split(string(convertedData), "\n")

	lines = lines[:len(lines)-1]

	var nonEmptyLines []string
	for _, line := range lines {
		// 行の前後の空白をトリムして、空でない場合だけ追加
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	dataAsBytes := []byte(strings.Join(nonEmptyLines, "\n"))

	// 変換後のデータを処理する
	err = DownloadAndSaveRaceSchedule(c, dataAsBytes)
	if err != nil {
		log.Println("レース情報の保存エラー:", err)
		return fmt.Errorf("レース情報の保存に失敗しました: %v", err)
	}

	return nil
}

func toHalfWidth(input string) string {
	return strings.Map(func(r rune) rune {
		// 全角数字・英字を半角に変換
		if r >= 'Ａ' && r <= 'Ｚ' {
			return r - 'Ａ' + 'A'
		}
		if r >= 'ａ' && r <= 'ｚ' {
			return r - 'ａ' + 'a'
		}
		if r >= '０' && r <= '９' {
			return r - '０' + '0'
		}
		if r == '　' { // 全角スペースを半角スペースに変換
			return ' '
		}
		if r == '：' { // 全角コロンを半角コロンに変換
			return ':'
		}
		return r
	}, input)
}
