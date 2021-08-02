package main

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

var (
	secretMapUID = sync.Map{}
)

func index(c *gin.Context) {
	c.String(200, "index")
}
func login(c *gin.Context) {
	time.Sleep(time.Millisecond * 320)
	auth(c)
	time.Sleep(time.Millisecond * 120)
	c.String(200, "login")

}
func userInfo(c *gin.Context) {
	c.String(200, "userInfo")

}
func jeagerMid(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(c.Request.URL.Path)
	span.Tracer().StartSpan("123", opentracing.ChildOf(span.Context()))
	defer span.Finish()
	SetTag(c, span, span.Context())
	c.Set(tracerKey, tracer)
	c.Set(spanCtxKey, span.Context())
	c.Next()
}
func runHttp(addr string) {
	engine := gin.New()
	engine.Use(jeagerMid)
	engine.GET("/", index)
	engine.GET("login", login)
	engine.GET("userInfo", userInfo)

	// run the server
	log.Fatal(engine.Run(addr))
}
func SetTag(c *gin.Context, span opentracing.Span, spanContext opentracing.SpanContext) {
	jaegerSpanContext := spanContextToJaegerContext(spanContext)
	span.SetTag("traceID", jaegerSpanContext.TraceID().String())
	span.SetTag("spanID", jaegerSpanContext.SpanID().String())
}

func spanContextToJaegerContext(spanContext opentracing.SpanContext) jaeger.SpanContext {
	if sc, ok := spanContext.(jaeger.SpanContext); ok {
		return sc
	} else {
		return jaeger.SpanContext{}
	}
}
