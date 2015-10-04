package config

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
	"log"
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

func LoadConfig() *Config {
	if config == nil {
		cfg := &Config{
			Env: *env,
		}

		cfg.DriverName = *driverName
		cfg.DataSourceName = *dataSourceName

		if *driverName == DefaultDriverName || *dataSourceName == DefaultDataSourceName {
			gDbConf, err := goose.NewDBConf("./db", *env, "")
			if err != nil {
				log.Fatal(err)
			}
			if *driverName == DefaultDriverName {
				cfg.DriverName = gDbConf.Driver.Name
			}
			if *dataSourceName == DefaultDataSourceName {
				cfg.DataSourceName = gDbConf.Driver.OpenStr
			}
		}

		config = cfg
	}

	return config
}
