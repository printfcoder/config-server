package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
)

var (
	db    *gorm.DB
	gRepo *repo
)

func NewGROMRepo(opts ...Option) *repo {
	if gRepo != nil {
		return gRepo
	}

	var options Options
	for _, o := range opts {
		o(&options)
	}
	db = newGORM(options)

	return &repo{}
}

func newGORM(options Options) *gorm.DB {
	var err error
	db, err = gorm.Open(options.Dialect, options.URL)
	if err != nil {
		log.Fatal(err)
	}

	db.SingularTable(options.SingularTable)
	db.DB().SetMaxOpenConns(options.MaxOpenConns)
	db.DB().SetMaxIdleConns(options.MaxIdleConns)
	db.LogMode(options.LogMode)

	if options.AutoMigrate {
		if err = db.AutoMigrate(options.AutoMigrateModels...).Error; err != nil {
			_ = db.Close()
			log.Fatal(err)
		}
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
