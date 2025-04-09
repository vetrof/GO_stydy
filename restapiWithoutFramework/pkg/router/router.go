package router

import (
	"context"
	"net/http"
)

type Context interface {
	context.Context
	SetCancellableCtx(ctx context.Context, cancel context.CancelFunc)
	SetHttpWriter(w http.ResponseWriter)
	SetHttpRequest(r *http.Request)
	Decode(dest interface{}) error
	WriteResponse(statusCode int, resp interface{}) error
}
