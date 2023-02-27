package gosdk

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/tracer"
	tr "go.opentelemetry.io/otel/trace"
	"time"
)

var tracerService tracer.Service

func SetUpTracer(conf *JaegerConfig) error {
	service, err := tracer.NewService(conf.Host, conf.Environment, conf.ServiceName)
	if err != nil {
		return err
	}

	tracerService = service

	return nil
}

func StartTracer(tracerName string, ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	if tracerService == nil {
		return nil, nil
	}

	return tracerService.Tracer(tracerName, ctx, name, opt...)
}

func CloseTracer() error {
	if tracerService == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return tracerService.Shutdown(ctx)
}
