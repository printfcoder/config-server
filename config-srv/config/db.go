package config

type DBConfig interface {
	GetDialect() string
	GetORM() string
	GetMysql() MysqlConfig
	GetGORM() GORMConfig
}

type dbConfig struct {
	Dialect string      `json:"dialect"`
	ORM     string      `json:"orm"`
	Mysql   mysqlConfig `json:"mysql"`
	GORM    gormConfig  `json:"gorm"`
}

func (d *dbConfig) GetDialect() string {
	return d.Dialect
}

func (d *dbConfig) GetORM() string {
	return d.ORM
}

func (d *dbConfig) GetMysql() MysqlConfig {
	return d.Mysql
}

func (d *dbConfig) GetGORM() GORMConfig {
	return d.GORM
}
