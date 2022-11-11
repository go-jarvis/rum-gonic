package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tangx/ginbinder"
)

type TraceParams struct {
	TraceId string `header:"traceId"`
	SpanId  string `header:"spanId"`
}

func (tp *TraceParams) SetDefault() {
	if tp.TraceId == "" {
		tp.TraceId = RandId()
	}

	if tp.SpanId == "" {
		tp.SpanId = RandId()
	}
}

// func Trace() rum.MiddlewareOperator {

// 	return rum.NewMiddleware(TraceHandler)
// }

func RandId() string {
	return uuid.New().String()
}

func TraceHandler(c *gin.Context) {
	cc := c.Copy()

	tp := &TraceParams{}
	_ = ginbinder.ShouldBindRequest(cc, tp)

	tp.SetDefault()
	spanId := RandId()
	c.Set("traceId", tp.TraceId)
	c.Set("pspanId", tp.SpanId)
	c.Set("spanId", spanId)

	c.Writer.Header().Set("traceId", tp.TraceId)
	c.Writer.Header().Set("spanId", tp.SpanId)
	c.Next()
}
