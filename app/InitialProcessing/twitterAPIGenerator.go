package initialprocessing

import (
	"github.com/ChimeraCoder/anaconda"
)

// GetTwitterAPI config  twitterのAPI取得用処理
func GetTwitterAPI(config Config) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.KEY.Consumer)
	anaconda.SetConsumerSecret(config.KEY.Secret)
	api := anaconda.NewTwitterApi(config.KEY.Token, config.KEY.TokenSecret)
	return api
}
