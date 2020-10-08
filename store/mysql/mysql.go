package mysql

import (
	"fmt"

	"github.com/evt/rest-api-example/config"
	"github.com/jinzhu/gorm"
)

// MySQL is a shortcut structure to a mysqldb DB handler
type MySQL struct {
	*gorm.DB
}

// Dial creates new database connection to postgres
func Dial(cfg *config.Config) (*MySQL, error) {
	if cfg.MysqlDB == "" {
		return nil, nil
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDB))
	if err != nil {
		return nil, err
	}
	// Print SQL statements in debug mode
	if cfg.LogLevel == "debug" {
		db.LogMode(true)
		//db.SetLogger(logger.Get())
	}
	return &MySQL{db}, nil
}
