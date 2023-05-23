package jager

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
	"io"
)

var (
	GlobalTracer *opentracing.Tracer
	jagerCloser  io.Closer
)

func NewDefaultJager() {
	cfg := config.Configuration{
		ServiceName: "fastweb",
		Disabled:    false,
		RPCMetrics:  true,
		Reporter: &config.ReporterConfig{
			CollectorEndpoint:          "",
			LocalAgentHostPort:         "",
			DisableAttemptReconnecting: false,
			LogSpans:                   true,
			User:                       "",
			Password:                   "",
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger), config.Metrics(metrics.NullFactory))
	if err != nil {

	}
	opentracing.SetGlobalTracer(tracer)
	GlobalTracer = &tracer
	jagerCloser = closer
}

func CloseJagerService() {
	jagerCloser.Close()
}
