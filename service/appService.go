package service

import (
	"dcat/models"
	"fmt"
	"log"
	"time"
)

// 发布帖子。入topics和topics_ex库
func Addapp(appinfo *models.AppInfo) (errMsg string, err error) {
	appinfo.Create_time = time.Now().Format("2006-01-02 15:04:05")
	_, temperr := models.NewAppDao().Insert(*appinfo)
	if temperr != nil {
		errMsg = "内部服务器错误"
		fmt.Println(errMsg, "：", temperr)
		return
	}

	return
}

func GetAppUserinfo(aid int) *models.AppList {

	applist := models.NewAppDao().FindAllByAid(aid)

	log.Println(applist.Total)
	for _, value := range applist.Rows {
		log.Println(value.AppName, value.PackageName)
	}
	return applist
}
