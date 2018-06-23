package initialprocessing

import (
	"os"
)

// FolderGenerasion フォルダがない場合生成する
func FolderGenerasion(dir string, useID string) {
	fInfo, _ := os.Stat(dir + useID)
	if fInfo == nil || !fInfo.IsDir() {
		os.Mkdir(dir+useID, 0777)
	}
}
