package middleware

import (
	"net/http"

	"github.com/ExcitingFrog/go-core-common/jaeger"
)

func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return Inject(
		handler,
	)
}

func Inject(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx, span, _ := jaeger.StartSpanAndLogFromContext(r.Context(), "Middleware:Inject")
		defer span.End()

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(f)
}
