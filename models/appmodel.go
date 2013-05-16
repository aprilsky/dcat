package models

type AppInfo struct {
	AppName     string `json:"appname"`
	Platform    string
	PackageName string `PK`
	Market      string
	Version     string
	Id          int
	Create_time string
	*Dao
}

type AppList struct {
	AppInfoList []AppInfo
}

func NewAppInfo() *AppInfo {
	return &AppInfo{
		Dao: &Dao{tablename: "applist"},
	}
}

func (this *AppInfo) prepareInsertData() {
	this.columns = []string{"app_name", "platform", "package_name", "market", "version", "create_time"}
	this.colValues = []interface{}{this.AppName, this.Platform, this.PackageName, this.Market, this.Version, this.Create_time}
}

func (this *AppInfo) Insert() (int, error) {
	this.prepareInsertData()
	result, err := this.Dao.Insert()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (this *AppInfo) GetList() (AppList, error) {
	err := this.ConnDb()
	if err != nil {
		return err
	}
	defer this.Close()
	strSql := "select * from applist"
	row := this.QueryRow(strSql, nil)

	return
}
