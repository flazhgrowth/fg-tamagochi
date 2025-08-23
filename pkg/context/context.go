package context

import "context"

type ContextKey struct{}

var (
	contextKeyUseMasterDB ContextKey = struct{}{}
)

// UseMasterDB sets the context to use master database
/*	By default, SQLator SQLReader uses the slave database.
	Use this function to set the context to use master database.
	Under the hood, SQLator SQLReader methods (Get and Find) will first checks through function contextlib.IsUseMasterDB. If it does, it then use the master database, instead of the slave database.
*/
func UseMasterDB(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyUseMasterDB, true)
}

// IsUseMasterDB checks whether the context is set to use master database or not
/*
	By default, SQLator SQLReader uses the slave database.
	Use this function to check whether the context is set to use master database or not.
	Note that you will not necessarily use this function directyly. You can use contextlib.UseMasterDB to set the context to use master database.
*/
func IsUseMasterDB(ctx context.Context) bool {
	val := ctx.Value(contextKeyUseMasterDB)
	if val == nil {
		return false
	}

	useMasterDB, ok := val.(bool)
	if !ok {
		return false
	}

	return useMasterDB
}
