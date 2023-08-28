package middleware

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return Inject(
		handler,
	)
}

type TraceparentHandler struct {
	next  http.Handler
	props propagation.TextMapPropagator
}

func Inject(h http.Handler) http.Handler {
	return &TraceparentHandler{
		next:  h,
		props: otel.GetTextMapPropagator(),
	}
}

func (h *TraceparentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := req.Header
	h.props.Inject(req.Context(), propagation.HeaderCarrier(w.Header()))
	remoteTraceIDBytes, _ := trace.TraceIDFromHex(header["X-B3-Traceid"][0])
	remoteSpanIDBytes, _ := trace.SpanIDFromHex(header["X-B3-Spanid"][0])

	parentSpanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: remoteTraceIDBytes,
		SpanID:  remoteSpanIDBytes,
	})

	req = req.WithContext(trace.ContextWithSpanContext(req.Context(), parentSpanContext))

	span := trace.SpanFromContext(req.Context())
	defer span.End()

	h.next.ServeHTTP(w, req)
}
