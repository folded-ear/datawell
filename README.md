# datawell

## install

Go 1.5 is required, and you have to enable the vendor experiment:

    $ go version
    go version go1.5.1 linux/amd64
    $ export GO15VENDOREXPERIMENT=1

You'll also need glide to get started:

    $ go get github.com/Masterminds/glide
    $ glide up
    [INFO] Fetching updates for ...
    ....
    $ glide rebuild
    [INFO] Building dependencies.
    ....
    
To run migrations, you'll need goose too:

    $ go get bitbucket.org/liamstask/goose/cmd/goose
    
