package gorm

type Options struct {
	Dialect           string
	SingularTable     bool
	URL               string
	MaxOpenConns      int
	MaxIdleConns      int
	LogMode           bool
	AutoMigrate       bool
	AutoMigrateModels []interface{}
}

type Option func(o *Options)

func WithDialect(dialect string) Option {
	return func(o *Options) {
		o.Dialect = dialect
	}
}

func WithSingularTable(singularTable bool) Option {
	return func(o *Options) {
		o.SingularTable = singularTable
	}
}

func WithURL(url string) Option {
	return func(o *Options) {
		o.URL = url
	}
}

func WithMaxOpenConns(maxOpenConns int) Option {
	return func(o *Options) {
		o.MaxOpenConns = maxOpenConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(o *Options) {
		o.MaxIdleConns = maxIdleConns
	}
}

func WithLogMode(logMode bool) Option {
	return func(o *Options) {
		o.LogMode = logMode
	}
}

func WithAutoMigrate(autoMigrate bool) Option {
	return func(o *Options) {
		o.AutoMigrate = autoMigrate
	}
}

func WithAutoMigrateModels(autoMigrateModels []interface{}) Option {
	return func(o *Options) {
		o.AutoMigrateModels = autoMigrateModels
	}
}
