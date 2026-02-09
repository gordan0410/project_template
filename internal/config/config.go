package config

type EnvConfig struct {
	SystemConfig   SystemConfig
	MasterDBConfig DatabaseConfig
	RedisConfig    RedisConfig
}

type SystemConfig struct {
	AppMode  string `env:"system_app_mode" envDefault:"production"`
	Host     string `env:"system_host" envDefault:"http://localhost"`
	HTTPPort int    `env:"system_http_port" envDefault:"9011"`
	ID       int    `env:"system_id" envDefault:"1"`
	LogLevel int    `env:"system_log_level" envDefault:"1"` // 0: debug, 1: info, 2: warn, 3: error 4:fatal 5:panic
}

type DatabaseConfig struct {
	Host         string `env:"master_db_host" envDefault:"localhost"`
	Port         int    `env:"master_db_port" envDefault:"3306"`
	DBName       string `env:"master_db_name" envDefault:"database"`
	Account      string `env:"master_db_account" envDefault:"root"`
	Password     string `env:"master_db_password" envDefault:"root"`
	MaxOpenConns int    `env:"master_max_open_connections" envDefault:"20"`
	MaxIdleConns int    `env:"master_max_idle_connections" envDefault:"20"`
	MaxConnLife  int    `env:"master_max_connection_lifetime" envDefault:"5"`
}

type RedisConfig struct {
	Host     string `env:"redis_host" envDefault:"localhost"`
	Port     int    `env:"redis_port" envDefault:"6379"`
	Password string `env:"redis_password" envDefault:"root"`
	DBNum    int    `env:"redis_db_num" envDefault:"0"`
	UseTLS   bool   `env:"redis_use_tls" envDefault:"false"`
	PoolSize int    `env:"redis_pool_size" envDefault:"1000"`
	MinIdle  int    `env:"redis_min_idle" envDefault:"50"`
}
