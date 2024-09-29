package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"wrpc.io/go/nats"
)

// extractTraceHeaderContext extracts the trace context from the wrpc headers.
func extractTraceHeaderContext(ctx context.Context) context.Context {
	headers, ok := wrpcnats.HeaderFromContext(ctx)
	if !ok {
		return ctx
	}

	pr := propagation.MapCarrier{}
	pr.Set("traceparent", headers.Get("traceparent"))
	pr.Set("tracestate", headers.Get("tracestate"))
	pr.Set("baggage", headers.Get("baggage"))

	return otel.GetTextMapPropagator().Extract(ctx, pr)
}
