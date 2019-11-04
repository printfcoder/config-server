package config

// MysqlConfig mysql 配置 接口
type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// mysqlConfig mysql 配置
type mysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

// URL mysql 连接
func (m mysqlConfig) GetURL() string {
	return m.URL
}

// Enabled 激活
func (m mysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m mysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m mysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
