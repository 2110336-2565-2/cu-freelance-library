package tracer

import (
	"context"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"go.opentelemetry.io/otel/sdk/trace"
	tr "go.opentelemetry.io/otel/trace"
)

func NewService(conf *gosdk.JaegerConfig) (Service, error) {
	tracerProvider, err := gosdk.InitJaegerTracerProvider(conf)
	if err != nil {
		return nil, err
	}

	return &tracer{
		tracerProvider: tracerProvider,
	}, nil
}

type tracer struct {
	tracerProvider *trace.TracerProvider
	tracer         tr.Tracer
}

func (t *tracer) Tracer(tracerName string) {
	tracer := t.tracerProvider.Tracer(tracerName)

	t.tracer = tracer
}

func (t *tracer) Start(ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	return t.tracer.Start(ctx, name, opt...)
}
