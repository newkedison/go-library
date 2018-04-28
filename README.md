# Library for golang (by newk)

This repo is a collection of my golang packages


### Install

go get -u github.com/newkedison/go-library


### For developer

* Use [dep](https://github.com/golang/dep) as dependency management tool 

    * Installation
    
        see https://golang.github.io/dep/docs/installation.html

    * Usage
    
        **dep init** :  Initialize the project
        
        **dep ensure** : Check all go file and add/remove package in vendor folder
        
        **dep ensure -add package1 [package2] ...**: Manual add package, **note** that if you want to add more than one package, you must use one command to add them, use multiple commands will only keep last package, see the [docs](https://golang.github.io/dep/docs/daily-dep.html#adding-a-new-dependency) for the reason

* Use [goconvey](http://goconvey.co/) for test

    * Installation and Usage
    
        $ go get github.com/smartystreets/goconvey
        
        $ $GOPATH/bin/goconvey
        
        Then open your browser to localhost:8080. Tests will be run from the working directory on down.
        
    * If you prefer **go test** , run as below in the root folder:
    
        go test ./...


### License

[BSD 3-Clause License](https://github.com/newkedison/go-library/blob/master/LICENSE)
