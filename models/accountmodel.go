package models

type AccountInfo struct {
	Aid         int
	AccountName string
	Pwd         string
	*Dao
}

func NewAccountInfo() *AccountInfo {
	return &AccountInfo{
		Dao: &Dao{tablename: "account"},
	}
}

func (this *AccountInfo) FindByAccountName(name string) error {
	err := this.ConnDb()
	if err != nil {
		return err
	}
	defer this.Close()
	strSql := "select * from account where account_name=?"
	row := this.QueryRow(strSql, name)

	var nameString string
	var aid int
	var pwd string

	err = row.Scan(&aid, &nameString, &pwd)

	this.Aid = aid
	this.Pwd = pwd
	this.AccountName = nameString

	return err
}
