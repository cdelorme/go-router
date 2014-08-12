
# go-router

Basic routing abstraction for the core http package in golang.

It aims at delivering a simple way to setup a server, and route traffic.


## toolkit alternatives

There are (at least) two decent feature-filled alternatives that include a variety of extras beyond just routing:

- [Gorilla](http://www.gorillatoolkit.org/)
- [gocraft](https://github.com/gocraft/web)


## framework alternatives

There are (at least) two complete web frameworks for go that will not only handle routing but the entire web stack:

- [martini](http://martini.codegangsta.io/)
- [revel](http://revel.github.io/)


## sales pitch

My router aims for simplicity while remaining completely modular.  It features the following:

- route creation and registration
- controller registration with standard-type compatible callbacks
- implements http.Handler interface (ServeHTTP) and can be applied directly to a server
- accepts a server by-reference or makes one on demand

You can use my router to register routes and start up a server, but won't need to reference it from controllers or other sections of code, keeping it as opaque as possible.

By default my library will ignore requests for `favicon`.  _My router has one external dependency; my logger package._


## usage

You can import my router with:

    import (
        "github.com/cdelorme/go-logger"
        "github.com/cdelorme/go-routing"
    )

You can create a route like this (it depends on my logger):

    logger := log.Logger{}
    router := routing.Router{Log: &log}

You can register single routes like this:

    router.CreateAndRegisterRoute("/path/", callbackOne, "GET")
    router.CreateAndRegisterRoute("/path2/", callbackTwo, "GET", "POST", "PUT")

_The callbacks can be part of another struct or stand-alone, and it must accept `http.ResponseWriter` and `http.Request` parameters._

You can register a compatible controller like this:

    router.RegisterController(controller, "prefix/")

_Your controller would be responsible for appending the prefix, which will be passed alongside the `CreateAndRegisterRoute` method, which can then be run from the controller._

You can apply my router to an existing server like this:

    server.Handler = router

Alternatively, my router can attach itself and start the server for you like this:

    router.Start(&server, "address:port")

_When you supply a server the address can be an empty string._

Finally, if you don't want to create a router and just want to spinup a new server from my router you can:

    router.Start(nil, "address:port")

_In this case the address is not optional._


## planned features

- allow registration of 404 handler
- remove Start() from the router (that logic doesn't really belong there)
- make the ignoreList a manageable property (to add and remove from)
- remove prefix from controller registration (not a reliable dependency, adds more unnecessary complexity)


# references

- [A good review article](http://corner.squareup.com/2014/05/evaluating-go-frameworks.html)
