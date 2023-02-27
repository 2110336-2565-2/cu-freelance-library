package gosdk

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
	tr "go.opentelemetry.io/otel/trace"
)

type Jaeger interface {
	Start(ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span)
}

func NewJaeger(tracerProvider *trace.TracerProvider) Jaeger {
	return &jaeger{
		tracerProvider: tracerProvider,
	}
}

type jaeger struct {
	tracer         tr.Tracer
	tracerProvider *trace.TracerProvider
}

func (j *jaeger) Start(ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	return j.tracer.Start(ctx, name, opt...)
}
