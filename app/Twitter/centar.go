package twitter

import (
	initialprocessing "TwitterRec/app/InitialProcessing"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

// Centar メイン
func Centar(config initialprocessing.Config, userID string, API *anaconda.TwitterApi, uniq map[string]bool) error {
	v := url.Values{}
	v.Set("count", "100")
	v.Set("screen_name", userID)
	// v.Set("max_id", "971343801238175744")

	var inputDBDATA = map[string]string{}

	tweets, err := API.GetUserTimeline(v)
	if err != nil {
		return err
	}

	// DBアクセスのための初期設定
	db, err := FirstProcess()
	if err != nil {
		return err
	}

	/** ツイート軍からツイート単体取得 */
	for _, tweet := range tweets {
		var mediaString = ""
		if tweet.Entities.Media != nil || len(tweet.Entities.Media) != 0 {
			/** ツイート内のメディア繰り返し */
			for _, onemedia := range tweet.ExtendedEntities.Media {
				// ファイルのURL取得
				fileURL := onemedia.Media_url_https
				pos := strings.LastIndex(fileURL, ".")
				err := download(config.MEDIA.Folder+userID+"/"+onemedia.Id_str+fileURL[pos:], fileURL)

				if err != nil {
					return err
				}

				// メディアのIDを結合
				mediaString += onemedia.Id_str
				mediaString += ","
			}
		}
		// メディアIDから余分な「,」を削除
		var outputMedia = ""
		if len(mediaString) != 0 {
			if mediaString != "" {
				sc := []rune(mediaString) // 文字数カウント
				outputMedia = string(sc[:(len(sc) - 1)])
			}
		}

		// 日付変換
		// layout := "2006-01-02 15:04:05"
		// time, _ := time.Parse(layout, tweet.CreatedAt)

		// ツイートのID
		inputDBDATA["ID"] = tweet.IdStr
		// ユーザ名
		inputDBDATA["USER"] = tweet.User.Name
		// ツイート内容
		inputDBDATA["TWEETTEXT"] = tweet.FullText
		// ツイートした日付
		inputDBDATA["DATE"] = tweet.CreatedAt
		// ツイートしたクライアント
		inputDBDATA["Source"] = tweet.Source
		// twitterのメディアのID(削除した文字列)
		inputDBDATA["mediaIDS"] = outputMedia

		if uniq[tweet.IdStr] != true {
			// DB登録
			inputSQL(db, inputDBDATA)
			uniq[tweet.IdStr] = true
		}
	}
	dbclose(db)
	return err
}

/** ダウンロード処理 */
func download(filename string, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if !Exist(filename) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		io.Copy(file, response.Body)
	}
	return err
}

// Exist ファイルが存在すればTRUE
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
