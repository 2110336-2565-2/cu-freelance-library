package tracer

import (
	"context"
	tr "go.opentelemetry.io/otel/trace"
)

type Service interface {
	Tracer(tracerName string, ctx context.Context, spanName string, opt ...tr.SpanStartOption) (context.Context, tr.Span)
}
