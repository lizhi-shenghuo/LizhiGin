package service

import (
	"database/sql"
	"LizhiGin/global"
	"LizhiGin/model"
	"LizhiGin/model/request"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//@description: 回写配置

func writeConfig(viper *viper.Viper, conf map[string]interface{}) error {
	for k, v := range conf {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@description: 创建数据库(mysql)
func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func initDB(InitDBFunction ...model.InitDBFunc) (err error) {
	for _, v := range InitDBFunction {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//@description: 创建数据库并初始化
func InitDB(conf request.InitDB) error {
	baseSetting := map[string]interface{}{
		"mysql.path": "",
		"mysql.db-name":  "",
		"mysql.username": "",
		"mysql.password": "",
		"mysql.config":   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == "" {
		conf.Port = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	fmt.Println(dsn)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}
	setting := map[string]interface{}{
		"mysql.path":     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		"mysql.db-name":  conf.DBName,
		"mysql.username": conf.UserName,
		"mysql.password": conf.Password,
		"mysql.config":   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	if err := writeConfig(global.LizhiViper, setting); err != nil {
		return err
	}
	m := global.LizhiConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	linkDns := m.Username + ":" + m.Password + "@tcp(" +m.Path + ")" + m.Dbname + "?" + m .Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns,
		DefaultStringSize:         161,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.LizhiDB = db
	}

	err := global.LizhiDB.AutoMigrate(
		
		)
}