package tracer

import (
	"context"
	tr "go.opentelemetry.io/otel/trace"
)

type Service interface {
	Tracer(tracerName string)
	Start(ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span)
}
