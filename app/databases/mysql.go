package databases

import (
	"fmt"
	"my-task-api/app/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *configs.AppConfig) *gorm.DB {
	// declare struct config & variable connectionString
	// username:password@tcp(hostdb:portdb)/db_name
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}

// db migration
func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Project{})
	DB.AutoMigrate(&Task{})
}
