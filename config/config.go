package config

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
)

type Config struct {
	Env            string
	DriverName     string
	DataSourceName string
}

const (
	DefaultDriverName     = "use db/dbconf.yml"
	DefaultDataSourceName = "use db/dbconf.yml"
)

var (
	config         *Config
	env            = flag.String("env", "development", "the DB/Goose environment to use")
	driverName     = flag.String("db-driver-name", DefaultDriverName, "The database driver name (the first param to sql.Open)")
	dataSourceName = flag.String("db-datasource-name", DefaultDataSourceName, "The datasource name (the second param to sql.Open)")
)

func LoadConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}

	config := Config{
		Env: *env,
	}

	config.DriverName = *driverName
	config.DataSourceName = *dataSourceName

	if *driverName == DefaultDriverName || *dataSourceName == DefaultDataSourceName {
		gDbConf, err := goose.NewDBConf("./db", *env, "")
		if err != nil {
			return nil, err
		}
		if *driverName == DefaultDriverName {
			config.DriverName = gDbConf.Driver.Name
		}
		if *dataSourceName == DefaultDataSourceName {
			config.DataSourceName = gDbConf.Driver.OpenStr
		}
	}

	return &config, nil
}
