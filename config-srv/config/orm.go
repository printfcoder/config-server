package config

type GORMConfig interface {
	GetSingularTable() bool
	GetLogMode() bool
	GetAutoMigrate() bool
}

type gormConfig struct {
	SingularTable bool `json:"singularTable"`
	LogMode       bool `json:"logMode"`
	AutoMigrate   bool `json:"autoMigrate"`
}

func (g gormConfig) GetSingularTable() bool {
	return g.SingularTable
}

func (g gormConfig) GetLogMode() bool {
	return g.LogMode
}

func (g gormConfig) GetAutoMigrate() bool {
	return g.AutoMigrate
}
