package models

import (
	"database/sql"
	"dcat/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sort"
	"strings"
)

const (
	dbName      = "root:root@tcp(127.0.0.1:3306)/dcat?charset=utf8"
	mysqlDriver = "mysql"
)

type Dao struct {
	*sql.DB
	// 构造sql语句相关
	tablename string
	where     string
	whereVal  []interface{} // where条件对应中字段对应的值
	limit     string
	order     string
	// 插入需要
	columns   []string      // 需要插入数据的字段
	colValues []interface{} // 需要插入字段对应的值
	// 查询需要
	selectCols string // 想要查询那些字段，接在SELECT之后的，默认为"*"
}

func NewDao(tablename string) *Dao {
	return &Dao{tablename: tablename}
}

func (this *Dao) ConnDb() (err error) {
	this.DB, err = sql.Open(mysqlDriver, dbName)
	if err != nil {
		fmt.Println("database initialize error : ", err.Error())
		return
	}
	fmt.Println("connect ok")
	return
}

func (this *Dao) Columns() []string {
	return this.columns
}

func (this *Dao) ColValues() []interface{} {
	return this.colValues
}

func (this *Dao) GetLimit() string {
	return this.limit
}

func (this *Dao) Tablename() string {
	return this.tablename
}

func (this *Dao) SelectCols() string {
	if this.selectCols == "" {
		return "*"
	}
	return this.selectCols
}

func (this *Dao) GetWhere() string {
	return this.where
}

func (this *Dao) GetOrder() string {
	return this.order
}

// Insert 插入数据
func (this *Dao) Insert() (sql.Result, error) {
	strSql := util.InsertSql(this)
	fmt.Println("Insert sql:", strSql)
	err := this.ConnDb()
	if err != nil {
		return nil, err
	}
	defer this.Close()
	stmt, err := this.Prepare(strSql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(this.ColValues()...)
}

// FindAll 查找多条数据
func (this *Dao) FindAll(selectCol ...string) (*sql.Rows, error) {
	sort.Sort(sort.StringSlice(selectCol))
	this.selectCols = "`" + strings.Join(selectCol, "`,`") + "`"
	strSql := util.SelectSql(this)
	fmt.Println("FindAll sql:", strSql)
	err := this.ConnDb()
	if err != nil {
		return nil, err
	}
	defer this.Close()
	return this.Query(strSql, this.whereVal...)
}

// Find 查找单条数据
// colFieldMap 数据库表中列对应go中对象的字段
func (this *Dao) Find(colFieldMap map[string]interface{}, selectCol ...string) error {
	colNum := len(selectCol)
	if colNum == 0 || (colNum == 1 && selectCol[0] == "*") {
		selectCol = util.MapKeys(colFieldMap)
	}
	sort.Sort(sort.StringSlice(selectCol))
	this.selectCols = "`" + strings.Join(selectCol, "`,`") + "`"
	strSql := util.SelectSql(this)
	log.Println("Find sql:", strSql)
	err := this.ConnDb()
	if err != nil {
		return err
	}
	defer this.Close()
	row := this.QueryRow(strSql, this.whereVal...)
	scanInterface := make([]interface{}, 0, colNum)
	for _, column := range selectCol {
		scanInterface = append(scanInterface, colFieldMap[column])
	}
	err = row.Scan(scanInterface...)
	if err == sql.ErrNoRows {
		log.Println("Find", strSql, ":no result ret")
		return nil
	}
	return err
}
