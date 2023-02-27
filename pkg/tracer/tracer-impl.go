package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	jg "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	tr "go.opentelemetry.io/otel/trace"
	"time"
)

func initJaegerTracerProvider(host string, environment string, serviceName string) (*tracesdk.TracerProvider, error) {
	exp, err := jg.New(jg.WithCollectorEndpoint(jg.WithEndpoint(host + "/api/traces")))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			attribute.String("environment", environment),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func NewService(host string, environment string, serviceName string) (Service, error) {
	tracerProvider, err := initJaegerTracerProvider(host, environment, serviceName)
	if err != nil {
		return nil, err
	}

	return &tracer{
		tracerProvider: tracerProvider,
	}, nil
}

type tracer struct {
	tracerProvider *trace.TracerProvider
}

func (t *tracer) Tracer(tracerName string, ctx context.Context, spanName string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	return t.tracerProvider.Tracer(tracerName).Start(ctx, spanName, opt...)
}

func (t *tracer) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return t.tracerProvider.Shutdown(ctx)
}
