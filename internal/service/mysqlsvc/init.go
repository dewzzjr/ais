package mysqlsvc

import (
	"log"

	"github.com/dewzzjr/ais/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(c config.Database) *gorm.DB {
	log.Printf("connecting db [%s]...\n", c.DSN)
	db := mysql.Open(c.DSN)
	gdb, err := gorm.Open(db, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gdb
}
