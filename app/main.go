// Click here and start typing.
package main

import (
	initialprocessing "TwitterRec/app/InitialProcessing"
	"TwitterRec/app/Twitter"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	twitterResident()
	wg.Wait()
}

func twitterResident() {
	config := initialprocessing.ConfigGeneration()
	userIds := config.USER.Ids
	dir := config.MEDIA.Folder
	wg.Add(1)
	go twitterProgram(userIds, dir, config)
}
func twitterProgram(userIds []string, dir string, config initialprocessing.Config) {
	// twitterで不具合が起きたりしたときに再接続する。
	api := initialprocessing.GetTwitterAPI(config)
	// 一度だけ実行したい処理用のFLAG
	var isOneProcess = false
	// DBアクセスをへらすためのMAP
	uniq := make(map[string]bool)
	// 真骨頂の無限ループ
	for {
		// 各種ユーザ情報をコンフィグから取得して個別にユーザのタイムラインを取得
		for _, useID := range userIds {
			if !isOneProcess {
				// ユーザ名を元にフォルダ生成
				initialprocessing.FolderGenerasion(dir, useID)

			}
			err := twitter.Centar(config, useID, api, uniq)

			// 問題が起きた際に30秒停止後再ログイン
			if err != nil {
				time.Sleep(30 * time.Second) // 3秒休む
				api = initialprocessing.GetTwitterAPI(config)
				isOneProcess = true
			}
		}
		if len(uniq) > 4000 {
			uniq = make(map[string]bool)
		}
		time.Sleep(5 * time.Second) // 5秒休む
	}

}
