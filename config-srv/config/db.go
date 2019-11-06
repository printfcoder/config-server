package config

type DBConfig interface {
	GetDialect() string
	GetMysql() MysqlConfig
}

type dbConfig struct {
	Dialect string      `json:"dialect"`
	Mysql   mysqlConfig `json:"mysql"`
}

func (d *dbConfig) GetDialect() string {
	return d.Dialect
}

func (d *dbConfig) GetMysql() MysqlConfig {
	return d.Mysql
}
