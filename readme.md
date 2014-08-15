
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
- customizable 404 (not found) error handler
- an ignore list for routes such as `/favicon.ico`

My router acts as an abstraction for registering many routes under a single `http.Handler` that can be assigned to an `http.Server`; without adding any references from the controllers to my library you can register all of your routes, keeping things as opaque as possible.


## usage

You can import my router (and its dependency) like this:

    import (
        "github.com/cdelorme/go-routing"
    )

You can create a router (with its dependency) like this:

    router := routing.Router{}

You can register single routes like this:

    router.CreateAndRegisterRoute("/path", callbackOne, "GET")
    router.CreateAndRegisterRoute("/path2", callbackTwo, "GET", "POST", "PUT")

_Callbacks can be stand-alone functions or part of a struct, both work provided they accept the right parameters.  Routes without a trailing slash will work with and without; this is to support various frontend services and browser specs.  If you do not want to support two paths, add the trailing slash to the route._

You can also register a custom 404 (NotFound) handler, like this:

    router.HandleNotFound = My404Handler

You can register a compatible controller like this:

    router.RegisterController(controller)

_Your controller would be responsible for appending the prefix, which will be passed alongside the `CreateAndRegisterRoute` method, which can then be run from the controller._

You can apply my router to an existing server like this:

    server.Handler = router

Or in more detail, like this:

    server := http.Server{
        Addr:           Address,
        MaxHeaderBytes: 1 << 20,
        Handler:        router,
    }


# references

- [A good review article](http://corner.squareup.com/2014/05/evaluating-go-frameworks.html)
