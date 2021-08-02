package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

const (
	tracerKey  = "tracer"
	spanCtxKey = "spanCtx"
)

func NewJaegerTracer(serviceName string) (opentracing.Tracer, error) {
	cfg := jaegerCfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: defaultJeagerAddr,
		},
		Tags: []opentracing.Tag{
			{
				Key:   "namespace",
				Value: "beta",
			},
		},
	}
	tracer, _, err := cfg.NewTracer(
		jaegerCfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		log.Println("tracer error ", err)
	}
	return tracer, err
}
func getInjectParent(c *gin.Context) (opentracing.Tracer, opentracing.SpanContext, bool) {
	var tracer opentracing.Tracer
	var spanCtx opentracing.SpanContext
	var ok bool
	tracerData, ok := c.Get(tracerKey)
	if !ok {
		return nil, nil, ok
	}
	tracer, ok = tracerData.(opentracing.Tracer)
	if !ok {
		return nil, nil, ok
	}
	spanCtxData, ok := c.Get(spanCtxKey)
	if !ok {
		return nil, nil, ok
	}
	spanCtx, ok = spanCtxData.(opentracing.SpanContext)
	return tracer, spanCtx, ok
}

func initJaeger() {
	tracer, _ := NewJaegerTracer("http.gin")
	opentracing.SetGlobalTracer(tracer)
}
