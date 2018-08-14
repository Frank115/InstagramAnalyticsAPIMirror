package mysql

import (
	"github.com/ProyectoIT/InstagramAnalyticsAPI/config"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/domain"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm dialects mysql
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open("mysql", config.DBUSER+":"+config.DBPASS+"@/"+config.DBNAME)
	if err != nil {
		return
	}
	DB.DropTableIfExists(&domain.User{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&domain.User{})
	return
}
