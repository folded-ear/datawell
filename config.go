package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
	"github.com/vharitonsky/iniflags"
)

type config struct {
	env            string
	driverName     string
	dataSourceName string
}

const (
	DEFAULT_DRIVER_NAME     = "use db/dbconf.yml"
	DEFAULT_DATASOURCE_NAME = "use db/dbconf.yml"
)

var (
	env            = flag.String("env", "development", "the DB/Goose environment to use")
	driverName     = flag.String("db-driver-name", DEFAULT_DRIVER_NAME, "The database driver name (the first param to sql.Open)")
	dataSourceName = flag.String("db-datasource-name", DEFAULT_DATASOURCE_NAME, "The datasource name (the second param to sql.Open)")
)

func newConfig() (*config, error) {
	iniflags.Parse()
	config := config{
		env: *env,
	}

	config.driverName = *driverName
	config.dataSourceName = *dataSourceName

	if *driverName == DEFAULT_DRIVER_NAME || *dataSourceName == DEFAULT_DATASOURCE_NAME {
		gDbConf, err := goose.NewDBConf("./db", *env, "")
		if err != nil {
			return nil, err
		}
		if *driverName == DEFAULT_DRIVER_NAME {
			config.driverName = gDbConf.Driver.Name
		}
		if *dataSourceName == DEFAULT_DATASOURCE_NAME {
			config.dataSourceName = gDbConf.Driver.OpenStr
		}
	}

	return &config, nil
}
