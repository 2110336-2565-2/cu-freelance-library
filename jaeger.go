package gosdk

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/tracer"
	tr "go.opentelemetry.io/otel/trace"
)

var tracerService tracer.Service

func SetUpTracer(conf *JaegerConfig, tracerName string) error {
	service, err := tracer.NewService(conf)
	if err != nil {
		return err
	}

	tracerService = service
	service.Tracer(tracerName)

	return nil
}

func StartTracer(ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	if tracerService == nil {
		return nil, nil
	}

	return tracerService.Start(ctx, name, opt...)
}
