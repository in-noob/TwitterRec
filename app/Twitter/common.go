package twitter

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type twitter struct {
	ID        int       `sql:"not null;type:int`
	USER      string    `sql:"not null;type:string`
	TWEETTEXT string    `sql:TWEETTEXT`
	DATE      time.Time `sql:"not null;type:date"`
	DATEINT   int       `sql:"not null;type:int`
	Source    string    `sql:"not null;type:string`
	mediaIDS  string    `sql:"not null;type:string`
}

// FirstProcess a
func FirstProcess() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root" // パスワードを設定しておく
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "local_database"
	CHARSET := "?charset=utf8mb4"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + CHARSET
	db, err := gorm.Open(DBMS, CONNECT)
	db.LogMode(false)

	// マイグレーション
	db.AutoMigrate(&twitter{})

	return db, err
}

func inputSQL(db *gorm.DB, inputDBDATA map[string]string) {
	// 構造体のインスタンス化
	twitterEx := twitter{}

	// 日付
	layout := "Mon Jan 2 15:04:05 +0000 2006"
	timeLainDate, _ := time.Parse(layout, inputDBDATA["DATE"])
	timeString := timeLainDate.Format("20060102150405")
	// 挿入したい情報を構造体に与える
	twitterEx.USER = inputDBDATA["USER"]
	twitterEx.TWEETTEXT = inputDBDATA["TWEETTEXT"]
	twitterEx.DATE = timeLainDate
	// twitterEx.DATEINT = timeInt
	twitterEx.DATEINT, _ = strconv.Atoi(timeString)
	twitterEx.Source = inputDBDATA["Source"]
	twitterEx.mediaIDS = inputDBDATA["mediaIDS"]
	twitterEx.ID, _ = strconv.Atoi(inputDBDATA["ID"])

	// INSERTを実行
	db.Create(&twitterEx)
}
func dbclose(db *gorm.DB) {
	defer db.Close()
}
