package infrastucture

import (
	"case3/pkg/xlogger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func dbSetup() {
	var err error

	l := gormLogger.Default.LogMode(gormLogger.Silent)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.UserName,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DSN,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: l,
	})

	if err != nil {
		xlogger.Logger.Err(err).Msgf("Failed to connect database : %s", err)
		return
	}

	db = conn

	fmt.Println("Database connected")
	return
}
