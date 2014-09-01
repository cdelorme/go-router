
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

- registration of routes using standard data types
- registration can be handed off to controllers matching our defined interface
- implements `http.Handler` interface (ServeHTTP) and can be applied directly to a server
- customizable 404 (not found) error handler
- an ignore list for routes such (ex. `/favicon.ico`)

My router acts as an abstraction for registering many routes under a single `http.Handler` that can be assigned to an `http.Server`; without any crossover object or data type references that are specific to my library.  This keeps things fully decoupled and as basic as possible.


## usage

You can import my router like this:

    import "github.com/cdelorme/go-routing"

You can create a router like this:

    router := routing.Router{}

Route registration accepts a path and one or more methods _(1)_:

    router.CreateAndRegisterRoute("/path", callbackOne, "GET")
    router.CreateAndRegisterRoute("/path2", callbackTwo, "GET", "POST", "PUT")

You can also register a custom 404 (NotFound) handler, like this:

    router.HandleNotFound = My404Handler

You can register a compatible controller like this _(2)_:

    router.RegisterController(controller)

You can create a new server and apply the router to it like this:

Or in more detail, like this:

    server := http.Server{
        Addr:           Address,
        MaxHeaderBytes: 1 << 20,
        Handler:        &router,
    }

1. Callbacks can be stand-alone functions or part of a struct.  Routes are wildcard matched using [`strings.Index`](http://golang.org/pkg/strings/#Index), adding built-in support for _optional_ trailing slashes, and custom enhancements such as URL-based argument parsing supporting Clean URL structures.

2. Your controller is responsible for appending any prefixes to a route during registration (`CreateAndRegisterRoute`).

_Because of wildcard routing, registering parent and child routes is ill-advised with this system._

I had originally toyed with the idea of passing optional prefixes through the registration process, but concluded that it created an unnecessary level of complication.


# references

- [A good review article](http://corner.squareup.com/2014/05/evaluating-go-frameworks.html)
