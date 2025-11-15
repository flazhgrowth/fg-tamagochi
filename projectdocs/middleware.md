# Middleware
Tamagochi utilize chi middleware under the hood. The only thing that differentiate chi and tamagochi middleware is how to use it on route

## The Method
Just like chi, tamagochi accepts function like with type func(next http.Handler) http.Handler. Technically, if you, let's say have a middleware package for chi, with type func(next http.Handler) http.Handler, you can use it without modifying anything at all. 

The only difference in tamagochi middleware, is how to use it in routes. As you can see in router method, method `Use` accept an array of `middleware.HTTPMiddleware`, in ellipsis. Type `middleware.HTTPMiddleware` is actually an alias to string. This works as identifier, to middleware handler that you made. 

If you see in app initialization, you can fill out field Middleware. The type is `map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler`. So you need to register all middlewares here and its identifier, in order to use it later on your routes. The identifier itself does not have to be `middleware.HTTPMiddleware` type, considering `middleware.HTTPMiddleware` is an alias to string. So you can fill out identifier here as you please. Of course, we recommend you to have constant collection, so you won't have inconsistency when using the middleware, because, in theory, let's say you registered a middleware with identifier `super_middleware`, then, this will work
```
rtr.Scope(func(r router.Router) {
    r.Use("super_middleware")
    // another routes in here
})
```

Of course, this practice is prone to error. In this case, grouping every type to a group of constant is a recommended approach. This way, there won't be any inconsistency around your codebase.