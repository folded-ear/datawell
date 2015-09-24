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
	DefaultDriverName     = "use db/dbconf.yml"
	DefaultDataSourceName = "use db/dbconf.yml"
)

var (
	env            = flag.String("env", "development", "the DB/Goose environment to use")
	driverName     = flag.String("db-driver-name", DefaultDriverName, "The database driver name (the first param to sql.Open)")
	dataSourceName = flag.String("db-datasource-name", DefaultDataSourceName, "The datasource name (the second param to sql.Open)")
)

func newConfig() (*config, error) {
	iniflags.Parse()
	config := config{
		env: *env,
	}

	config.driverName = *driverName
	config.dataSourceName = *dataSourceName

	if *driverName == DefaultDriverName || *dataSourceName == DefaultDataSourceName {
		gDbConf, err := goose.NewDBConf("./db", *env, "")
		if err != nil {
			return nil, err
		}
		if *driverName == DefaultDriverName {
			config.driverName = gDbConf.Driver.Name
		}
		if *dataSourceName == DefaultDataSourceName {
			config.dataSourceName = gDbConf.Driver.OpenStr
		}
	}

	return &config, nil
}
