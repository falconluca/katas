package models

import (
	"blogapi/pkg/setting"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model 数据模型基础字段
type Model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// init 初始化数据库连接
func init() {
	var (
		err         error
		dbType      string
		dbName      string
		user        string
		password    string
		host        string
		tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("加载区失败 err: %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true) // 自动建表的表名称设置成单数
	db.LogMode(true)

	// 注册gorm全局钩子
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampForCreateCallback) // TODO callbackName有啥用?
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 配置数据源
	pool := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
}

func updateTimestampForCreateCallback(s *gorm.Scope) {
	if s.HasError() {
		return
	}

	n := time.Now().Unix()
	if createdOn, ok := s.FieldByName("CreatedOn"); ok && createdOn.IsBlank {
		createdOn.Set(n)
	}
	if modifiedOn, ok := s.FieldByName("ModifiedOn"); ok && modifiedOn.IsBlank {
		modifiedOn.Set(n)
	}
}

func updateTimestampForUpdateCallback(s *gorm.Scope) {
	if _, ok := s.Get("gorm:update_column"); ok {
		s.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(s *gorm.Scope) { // TODO debug 看看s里面的内容
	if s.HasError() {
		return
	}

	var opt string
	if str, ok := s.Get("gorm:delete_option"); ok {
		opt = fmt.Sprint(str) // 这是类型转换吗? 我可以这样处理吗 --> opt = str.(string)
	}

	var sql string
	deletedOn, ok := s.FieldByName("DeletedOn")
	if ok && !s.Search.Unscoped {
		// 存在deletedOn字段并且不是Unscoped时, 执行软删除 Soft Delete
		sql = fmt.Sprintf("update %v set %v=%v%v%v",
			s.QuotedTableName(),
			s.Quote(deletedOn.DBName),
			s.AddToVars(time.Now().Unix()), // 防止SQL注入
			addSpace2Str(s.CombinedConditionSql()),
			addSpace2Str(opt),
		)
	} else {
		// 硬删除
		sql = fmt.Sprintf("delete from %v%v%v",
			s.QuotedTableName(),
			addSpace2Str(s.CombinedConditionSql()),
			addSpace2Str(opt),
		)
	}
	s.Raw(sql).Exec()
}

func addSpace2Str(str string) interface{} {
	if str == "" {
		return ""
	}

	return " " + str
}

// CloseDb 关闭数据库连接
func CloseDb() {
	defer db.Close()
}
