package autoload

import "time"

type MysqlConfig struct {
	Enable       bool          `ini:"enable" yaml:"enable"`
	Host         string        `ini:"host" yaml:"host"`
	Username     string        `ini:"username" yaml:"username"`
	Password     string        `ini:"password" yaml:"password"`
	Port         uint16        `ini:"port" yaml:"port"`
	Database     string        `ini:"database" yaml:"database"`
	Charset      string        `ini:"charset" yaml:"charset"`
	TablePrefix  string        `ini:"table_prefix" yaml:"table_prefix"`
	MaxIdleConns int           `ini:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int           `ini:"max_open_conns" yaml:"max_open_conns"`
	MaxLifetime  time.Duration `ini:"max_lifetime" yaml:"max_lifetime"`
	LogLevel     int           `ini:"log_level" yaml:"log_level"`
	PrintSql     bool          `ini:"print_sql" yaml:"print_sql"`
}

var Mysql = MysqlConfig{
	Enable:       true,
	Host:         "49.233.42.135",
	Username:     "root",
	Password:     "blq200210@",
	Port:         3306,
	Database:     "test",
	Charset:      "utf8mb4",
	TablePrefix:  "",
	MaxIdleConns: 10,
	MaxOpenConns: 100,
	MaxLifetime:  time.Hour,
	LogLevel:     4,
	PrintSql:     true,
}
