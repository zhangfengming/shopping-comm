package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

/*
*
jaeger 的实例化方法
*/
func NewTracer(serviceName string, address string) (opentracing.Tracer, io.Closer, error) {
	config := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  address,
		},
	}

	return config.NewTracer()
}
