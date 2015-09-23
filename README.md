# datawell

![Build Status](https://travis-ci.org/folded-ear/datawell.svg)

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
core SQL functionality is used. In fact, other databases probably will work
too. Only PostgreSQL 9.2 is supported though, because I'm lazy and don't want
to worry about cross-platform stuff.  If _you_ care, however, I would be more
than happy to discuss, answer questions, and merge pull requests as needed.
    
## configuration

Configuration is provided by command-line flags:

    $ datawell --help
    Usage of datawell:
      -config string
    	Path to ini config for using in go flags. May be relative to the current executable path.
      -configUpdateInterval duration
            Update interval for re-reading config file set via -config flag. Zero disables config file re-reading.
      -db-datasource-name string
            The datasource name (the second param to sql.Open) (default "use-dbconf.yml")
      -db-driver-name string
            The database driver name (the first param to sql.Open) (default "use-dbconf.yml")
      -dumpflags
            Dumps values for all flags defined in the app into stdout in ini-compatible syntax and terminates the app.
      -env string
            the DB/Goose environment to use (default "development")

All options may also be provided via an INI file as specified by the
`--config` flag.  Flags override the config file values.  You can generate a
sample config file (to STDOUT) by using the `--dumpflags` flag:

    $ datawell --dumpflags
    configUpdateInterval = 0  # Update interval for re-reading config file set via -config flag. Zero disables config file re-reading.
    db-datasource-name = use db/dbconf.yml  # The datasource name (the second param to sql.Open)
    db-driver-name = use db/dbconf.yml  # The database driver name (the first param to sql.Open)
    env = development  # the DB/Goose environment to use

Further, the database connect information (driver name and open string) can
also be provided by a `./db/dbconf.yml` file according to Goose's specs.  This
has the lowest priority and will be superceded by both the INI file and flags,
but is useful for keeping DRY in development environments, and is _mandatory_
if you wish to run the Goose migrations.

Once configured, you can use Goose to bootstrap your database with the
necessary schema objects:

    $ goose up
