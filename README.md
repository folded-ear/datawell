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

