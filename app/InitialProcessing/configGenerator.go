package initialprocessing

import (
	"TwitterRec/app/Constant"
	"os"

	"github.com/BurntSushi/toml"
)

// Config TOML形式の設定ファイルから取得したものを定義
type Config struct {
	KEY   KeyrConfig
	MEDIA mediaConfig
	USER  userConfig
}

// KeyrConfig TOMLから取得した設定ファイル内の鍵に関する情報を内包
type KeyrConfig struct {
	Consumer    string
	Secret      string
	Token       string
	TokenSecret string
}

type mediaConfig struct {
	Folder string
}

type userConfig struct {
	Ids []string
}

// ConfigGeneration コンフィグ作成
func ConfigGeneration() Config {

	var config Config
	_, err := toml.DecodeFile(constant.ConfigPath+"twitter.tml", &config)
	if err != nil {
		panic(err)
	}
	// 設定ファイルに記載されている定数がなければ作成
	Folder(config)
	return config
}

// Folder ダウンロードフォルダ作成
func Folder(config Config) {

	fInfo, _ := os.Stat(config.MEDIA.Folder)
	if fInfo != nil {
		os.Mkdir(config.MEDIA.Folder, 0777)
	}
}
