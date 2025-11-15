# Routing and HTTP Handlers in Tamagochi
Tamagochi utilize chi routing under the hood. But Tamagochi wraps it under its method, that still familiar to chi, but also a bit different. Let's take a look.

## The Methods
Just like Chi, as we wraps chi method under its own methods, these HTTP Methods are available to use
```
type Router interface {
	Get(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Post(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Put(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Patch(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Delete(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Options(pattern string, h handler.HTTPHandlerFunc)

	Use(handlersNames ...fgmw.HTTPMiddleware)

	Group(pattern string, fn func(r Router)) Router

	Scope(fn func(r Router))

	Mount(pattern string, fn http.Handler)

	ServeDocs()

	ServeProfiler(pattern ...string)

	Routes() *chi.Mux
}
```
### [1] Initializing New Router
You can initialize a new Router using the Tamagochi `github.com/flazhgrowth/fg-tamagochi/pkg/http/router` package. Take a look at example below.
```
rtr := router.NewRouter()
```
Function `NewRouter` returns a Router interface mentioned above. Through this interface, you can use mentioned methods for routing

### [2] handler.HTTPHandlerFunc
The HTTPHandlerFunc from package `github.com/flazhgrowth/fg-tamagochi/pkg/http/handler` is a `func(r request.Request, w response.Response)`
Below is the example of how the method implementations looks like:
```
func (api *API) Lala(r request.Request, w response.Response) {
	ctx := r.GetContext()
	accountInfo, err := r.GetAccountInfo()
	if err != nil {
		w.Respond(nil, err)
		return
	}
	resp, err := api.accountUsecase.Me(ctx, accountInfo)
	if err != nil {
		w.Respond(nil, err)
		return
	}

	w.Respond(resp, nil)
}
```
Explanation:
1. **request.Request is from package `github.com/flazhgrowth/fg-tamagochi/pkg/http/request`**

As you can see. There are a few methods that you can use through `request.Request`. Here are the full interface definition of `request.Request`
```
type Request interface {
    // GetContext get request context
    GetContext() context.Context

    // DecodeBody decode body bytes to struct. The argument of this method accepts struct pointer
    DecodeBody(dest any) error

    // DecodeQueryParam decode request param to struct. The argument of this method accepts struct pointer
    DecodeQueryParam(dest any) error

    // GeneralHeaders gets general headers. This includes Host, UserAgent, Accept, AcceptEncoding, Referer, Connection
    GeneralHeaders() HTTPGeneralHeaders

    /*
        SecurityHeaders gets security headers. This includes Authorization, ProxyAuthorization, Cookie. There's IsAuth also.
        This is a helper to decide if Authorization header is available and valid (valid would be the value starts with prefix "Bearer").

        Authorization value also does not include the "Bearer" prefix anymore
    */
    SecurityHeaders() HTTPSecurityHeaders

    ContentHeaders() HTTPContentRelatedHeaders

    // GetAccountInfo gets account info from the context. Will return apierrors.ErrorUnauthorized() on failed get
    GetAccountInfo() (accountInfo entity.AccountInfo, err error)

    // GetNetHTTPHeaders gets native HTTP Headers
    GetNetHTTPHeaders() http.Header

    // NativeRequest returns the *http.Request native
    NativeRequest() *http.Request

    // URLParam gets url param value based on given key
    URLParam(key string) ParamsValue

    // URLParamDecode decode url param into struct with tag of urlparam and urlparamtype for its type
    /*
        Please note that, it is still recommended to use URLParam method instead of using URLParamDecode, due to efficiency
    */
    URLParamDecode(dest any) error
}
```
You will find yourself use `GetContext`, `DecodeBody`, and `DecodeQueryParam` frequently. 

`GetContext` will return context. 

`DecodeBody` decode the body of the request into a provided struct on the second argument, while `DecodeQueryParam` decodes query param into struct in which this struct fields needs to be annotated using `schema`, unlike `DecodeBody` that needs `json` annotation on its fields.

If you want to access query param, a bit more native on how we usually do it in chi, Tamagochi provide `NativeRequest` that will return `*http.Request` (from net/http). By that, you can access any method it has like how we normally would. Eg:
```
r.NativeRequest().URL.Query().Get("key") // equivalent to r.URL.Query().Get(""), assuming that r type is *http.Request
```

Now, we also have this `DecodeURLParam`. Like query param, Tamagochi provide "decode" method for URL param. It accept struct with a bit specialized annotation on its fields. Let's take a look at the sample below
```
AmazingStruct struct {
    Name            string `urlparam:"name" urlparamtype:"string"`
    SomeNumber      string `urlparam:"some_number" urlparamtype:"number"`
    SomeBoolean     bool   `urlparam:"some_boolean" urlparamtype:"bool"`
}
```
By this example, we assume you have a route that basically looks like this (eg: /api/v1/name/{name}/numba/{some_numba}/{some_boolean}).

Please note that the annotation also provide `urlparamtype` with values of string, number, and bool. Besides string and bool that are pretty straighforward, number will cast the value into int64. We recommend fraction value to be annotated using string, and you manually cast it yourself to your desired destination type.

You can always use `URLParam`, and access it one by one (and we still recommend you to use `URLParam` like how you normally would. `URLParamDecode` uses package reflect under the hood. So it will slower regardless. By a fraction most likely, but we still need to state this.)

2. **response.Response is from package `github.com/flazhgrowth/fg-tamagochi/pkg/http/response`**
There's not much we can do with `response.Response`. We use response to write HTTP response. If you see the above example, we use `w` (of type response.Response) method, which is `Respond`. Method `Respond` accepts 3 arguments, which is data (any) and error, and status code (optional).

By default, you won't need to pass status code in here. If the data passed is not nil, it will be Status Code 200 by default. If error passed, it will check if the error is an `github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors` type. If it's not, it will then use Status Code 500 by default.

But, maybe it's an insert/create api, in which, on success, you want to return Status Code 201. This is perfectly doable using the third argument like so:
```
w.Respond(data, nil, http.StatusCreated)
```

### [3] Router docs
Method Get, Post, Put, Patch, and Delete have ...RouterDocs as its third argument. This struct can be use to expose the route to docs. We use Scalar Docs for this. 

Take a look at RouterDocs struct below
```
RouterDocs struct {
    Security     SecAuth
    Request      any
    Response     any
    IsDeprecated bool
    Tags         string
    Summary      string
    Description  string
}
```
Every field in the struct is pretty straightforward, in term of usage. But to be more precise on these 3 fields:
1. Security type is router.SecAuth. There are 2 types of SecAuth as for now, router.SecurityAPIKey and router.SecurityBearerAuth
2. Request can be filled with request struct for HTTP method that accept body. While url param and query param, is also needed to be in struct form for the docs as well. Field for url param annotated using `path:"path_name"` while query param annotated using `query:"page"`
3. Response can be filled with response struct. There is no need to pass the BaseResponse. Tamagochi already handled it to also include base response, hence the struct passed will be place in field `Data` in BaseResponse.

### [4] Method Get(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)
Use this to route a GET method to an endpoint. Example:
```
rtr.Get("/me", api.Me)
```

### [5] Post(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)
Use this to route a POST method to an endpoint. Example:
```
rtr.Post("/todo", api.CreateTodo)
```

### [6] Put(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)
Use this to route a PUT method to an endpoint. Example:
```
rtr.Put("/todo/{id}", api.EditTodo)
```

### [7] Patch(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)
Use this to route a PATCH method to an endpoint. Example:
```
rtr.Patch("/todo/{id}", api.EditTodo)
```

### [8] Delete(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)
Use this to route a DELETE method to an endpoint. Example:
```
rtr.Patch("/todo/{id}", api.EditTodo)
```

### [9] Use(handlersNames ...middleware.HTTPMiddleware)
Use this method to attach a middleware to a route
```
rtr.Use(middleware.MIDDLEWARE_BASIC_BEARER_AUTH)
```
There are a few things to note with `Use` method. Normally in chi, method `Use` accepts a function that accepts http.Handler and returns http.Handler.

Tamagochi wraps this and instead of directly passing the function, method `Use` in Tamagochi accepts identifier, in which the type is middleware.HTTPMiddleware (`github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware`), and the identifier is related to the handler via `map[middleware.HTTPMiddleware]func(http.Handler) http.Handler`. 

This approach ensures that every needed middleware is registered on the main `Conjure` method (usually at `main.go`, and we recommend you to). Tamagochi also has a few middlewares already registered beforehand. This approach also makes any middleware you created can be use, using its identifier like so:
```
// assuming that you already registered a new middleware called "super_printer"
// you can use the registered middleware like so:
rtr.Use("super_printer") // type middleware.HTTPMiddleware is an alias to string. In which, its perfectly okay to pass string here. But we recommend you to make a constant with type middleware.HTTPMiddleware, so you can avoid of mistyping errors.
```
More of middlewares, please refer to [Middleware Section](./middleware.md)

### [10] Group(pattern string, fn func(r Router)) Router
Use Group when you want to group a few endpoints under its prefix pattern, like so:
```
parentR.Group("/resources", func(r router.Router) {
    r.Use("apikey")
    r.Post("/foo/bar", api.FooBar, nil)
    r.Get("/foo/bezt", api.FooBezt, nil)
})
```
This will then group `/foo/bar` and `/foo/bezt` under `/resources` endpoint. Hence to access this API, the full endpoint would be `/resources/foo/bar` and `/resources/foo/bezt`.

### [11] Scope(fn func(r Router))
Pretty much the same with Group, but instead of group a few endpoints under a prefix pattern, it just scope it without any prefix pattern. Useful if you don't want to use group, but you need to do something to multiple endpoints, lets say, a middleware.

### [12] Mount(pattern string, fn http.Handler)
Let's say you want to do something that a bit more native to chi, we provide Mount method that accepts pattern and native http.Handler.

### [13] ServeDocs(pattern ...string)
ServeDocs will serve the Scalar Docs created when you run the app. For now, our only option for security is to expose this on non production environment only. Of course, you can expose this manually like how you normally do. It will also give you a better control on how you want to serve the swagger.

### [13] ServeProfiler(pattern ...string)
Use this method if you want to serve profiler (using pprof under chi middleware Profiler). By default, it will be mounted under endpoint `/pprof/profiler`. But if you pass a pattern on the argument, it will then mount it through the pattern passed.