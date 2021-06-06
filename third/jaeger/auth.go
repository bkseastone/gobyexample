package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func auth(c *gin.Context) {
	tracer, parentSpanCtx, ok := getInjectParent(c)
	jaegerSpanContext := spanContextToJaegerContext(parentSpanCtx)
	// trace, _ := NewJaegerTracer("auth")
	spanName := "认证系统"
	log.Println(jaegerSpanContext)
	var span opentracing.Span
	if !ok {
		span = tracer.StartSpan(spanName)
	} else {
		span = tracer.StartSpan(spanName,
			opentracing.ChildOf(jaegerSpanContext),
			opentracing.Tag{Key: string(ext.Component), Value: "http component"},
		)
	}
	defer span.Finish()
	time.Sleep(time.Millisecond * 787)
}
