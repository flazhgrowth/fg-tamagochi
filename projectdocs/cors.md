# CORS
CORS in tamagochi by default use what chi use also. In tamagochi, it is configurable.

On app initialization, the config accepts `middleware.CorsOpt` on its cors field. Type `middleware.CorsOpt` has 2 fields, `Opts` and `ValidatorHandlers`. 

Field `Opts` type is `cors.Options` from `github.com/go-chi/cors` package. You can fill out the fields accordingly. Field `ValidatorHandlers` type is `[]middleware.FnCorsAdditionValidator`

Type `middleware.FnCorsAdditionValidator` is an alias to `func(request.Request) error`. Basically, you can add more validation to your liking, by adding more function like `func(request.Request) error`. Tamagochi will then run every handler, and return error if any of that handler return error. Of course, you can return any type error on this. But it will result in http status code 500. It is recommended to use apierrors errors, so you can control the status code.