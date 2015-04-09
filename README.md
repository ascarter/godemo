# godemo

A brief tour of the Go language

## Getting Started

Install Go:

* [Binary Distribution](http://golang.org/doc/install)
* Mac OS X: `brew install go`
* Ubuntu: `sudo apt-get install golang`


Create a workspace:

	$ mkdir $HOME/go
	$ export GOPATH=$HOME/go
	$ export PATH=$PATH:$GOPATH/bin

## Download Demo

    $ go get -u github.com/ascarter/godemo
    $ cd $GOPATH/src/github.com/ascarter/godemo

## Presentation

To view the presentation slides:

    go get golang.org/x/tools/cmd/present
    present -play=true

In browser navigate to `http://localhost:3999` to view slides
