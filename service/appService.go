package service

import (
	"dcat/models"
	"fmt"
	"time"
)

// 发布帖子。入topics和topics_ex库
func Addapp(appinfo *models.AppInfo) (errMsg string, err error) {
	appinfo.Create_time = time.Now().Format("2006-01-02 15:04:05")
	_, temperr := appinfo.Insert()
	if temperr != nil {
		errMsg = "内部服务器错误"
		fmt.Println(errMsg, "：", temperr)
		return
	}

	return
}

func GetAppUserinfo() {

}
