package models

import (
	"log"
)

type AppInfo struct {
	AppName     string `json:"appname"`
	Platforms   string `json:"platforms"`
	PackageName string `json:"packagename"`
	Market      string `json:"market"`
	Versions    string `json:"versions"`
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Create_time string `json:"createtime"`
}

type AppDao struct {
	*Dao
	AppInfo
}

type AppList struct {
	Total int        `json:"total"`
	Rows  []*AppInfo `json:"rows"`
}

func NewAppDao() *AppDao {
	return &AppDao{
		Dao: &Dao{tablename: "applist"},
	}
}

func (this *AppDao) Insert(app AppInfo) (int64, error) {
	err := this.ConnDb()
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}

	defer this.Close()

	stmt, err := this.Prepare("insert into applist(app_name, platform, package_name, market, version, create_time, aid) values(?,?,?,?,?,?,?)")

	defer stmt.Close()

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	result, err := stmt.Exec(app.AppName, app.Platforms, app.PackageName, app.Market, app.Versions, app.Create_time, app.Aid)
	line, err := result.RowsAffected()
	return line, err
}

func (this *AppDao) FindAllByAid(aid int) *AppList {

	err := this.ConnDb()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer this.Close()
	strSql := "select * from applist where aid=?"
	log.Println(strSql)
	rows, err := this.Query(strSql, aid)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	appList := new(AppList)

	for rows.Next() {
		app := new(AppInfo)
		if err := rows.Scan(&app.AppName, &app.Platforms, &app.PackageName, &app.Id, &app.Market, &app.Versions, &app.Create_time, &app.Aid); err != nil {
			log.Println(err)
		} else {
			appList.Rows = append(appList.Rows, app)
		}
	}
	appList.Total = len(appList.Rows)
	return appList
}

func (this *AppDao) FindAllByAccountName(name string) *AppList {

	err := this.ConnDb()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer this.Close()
	strSql := "select * from applist where aid=(select aid from account where account_name=?)"
	log.Println(strSql)
	rows, err := this.Query(strSql, name)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	appList := new(AppList)

	for rows.Next() {
		app := new(AppInfo)
		if err := rows.Scan(&app.AppName, &app.Platforms, &app.PackageName, &app.Id, &app.Market, &app.Versions, &app.Create_time, &app.Aid); err != nil {
			log.Println(err)
		} else {
			appList.Rows = append(appList.Rows, app)
		}
	}
	appList.Total = len(appList.Rows)
	return appList
}
