package gormdao

import (
	"fmt"
	"os"
	"time"

	"github.com/CRL-Studio/AuthServer/src/utils/config"
	_ "github.com/go-sql-driver/mysql" // init mysql driver

	"github.com/jinzhu/gorm"
)

var (
	db       *gorm.DB
	interval = config.Get("db.interval").(int)
	dialect  = config.Get("db.dialect").(string)
	source   = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Get("db.user"),
		config.Get("db.password"),
		config.Get("db.host"),
		config.Get("db.port"),
		config.Get("db.database"),
		config.Get("db.flag"),
	)
)

func init() {
	go connectionPool()
}

// DB return database instance
func DB() *gorm.DB {
	if db == nil {
		connect(dialect, source)
	}
	return db
}

// Close close database connection
func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
		}
	}
}

func connectionPool() {
	for {
		if err := DB().DB().Ping(); err != nil {
			connect(dialect, source)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func connect(dialect string, source string) {
	conn, err := gorm.Open(dialect, source)

	if err != nil {
		os.Exit(1)
	} else {
		conn.SingularTable(true)
		conn.BlockGlobalUpdate(true)
		conn.Exec("SET wait_timeout=300")

		db = conn
	}
}
