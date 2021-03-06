# datawell

[![Build Status](https://travis-ci.org/folded-ear/datawell.svg)](https://travis-ci.org/folded-ear/datawell)
[![Code Climate](https://codeclimate.com/github/folded-ear/datawell/badges/gpa.svg)](https://codeclimate.com/github/folded-ear/datawell)

## install

Go 1.5 is required, you have to enable the vendor experiment, and you'll want
to have your `bin` directory on your path:

    $ go version
    go version go1.5.1 linux/amd64
    $ export GO15VENDOREXPERIMENT=1
    $ export PATH=$GOPATH/bin:$PATH

You'll also need godep to get started:

    $ go get github.com/tools/godep
    $ godep restore

Now build it (this will also build all the dependencies, so it'll take a bit):
    
    $ go install

To run migrations, you'll need [Goose](https://bitbucket.org/liamstask/goose/):

    $ go get bitbucket.org/liamstask/goose/cmd/goose

## database

PostgreSQL 9.2 is "required". Other versions will probably work fine, as only
core SQL functionality is used. In fact, other databases would very likely also
work, though Go's SQL drivers do no have a consistent approach to bind
parameters, so it'd take a bit of effort to make it work

To this point there has not been a compelling reason to undertake that effort,
but if you have one (or even better, are willing to put in the time) we'd be
happy to discuss and help in whatever way we can.  The `sqlx` package can likely
provide a way around bind parameter differences, for example.
    
## configuration

Configuration is provided by command-line flags:

    $ datawell --help
    datawell is what EventLog became.
    
    Usage:
        datawell [options] <subcommand> [subcommand options]
    
    Options:
      -db-datasource-name string
            The datasource name (the second param to sql.Open) (default "use db/dbconf.yml")
      -db-driver-name string
            The database driver name (the first param to sql.Open) (default "use db/dbconf.yml")
      -env string
            the DB/Goose environment to use (default "development")
    
    Commands:
    greet      say hello
    
    serve      start the web server
      -port uint
            Port to listen on (default 8080)
    
    demo       run the demo

    echo       echo out flags and args

The database connect information (driver name and open string) can also be
provided by a `./db/dbconf.yml` file according to Goose's specs.  This will be
superceded by flags, but is useful for keeping DRY in development environments,
and is _mandatory_ if you wish to run the Goose migrations.

Once configured, you can use Goose to bootstrap your database with the
necessary schema objects:

    $ goose up
    ...
    $ goose status
    goose: status for environment 'development'
        Applied At                  Migration
        =======================================
        Sun Oct  4 21:37:08 2015 -- 20150923145448_Bootstrap.sql
