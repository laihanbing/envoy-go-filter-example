package main

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
	"google.golang.org/protobuf/types/known/anypb"

	xds "github.com/cncf/xds/go/xds/type/v3"
)

func init() {
	http.RegisterHttpFilterConfigFactoryAndParser("basic-auth", configFactory, &parser{})
}

type config struct {
	username string
	password string
}

type parser struct {
}

func (p *parser) Parse(any *anypb.Any, callback api.ConfigCallbackHandler) (interface{}, error) {
	configStruct := &xds.TypedStruct{}
	if err := any.UnmarshalTo(configStruct); err != nil {
		return nil, nil
	}

	v := configStruct.Value
	conf := &config{}
	if username, ok := v.AsMap()["username"].(string); ok {
		conf.username = username
	}
	if password, ok := v.AsMap()["password"].(string); ok {
		conf.password = password
	}
	callback.DefineCounterMetric("test")
	callback.DefineGaugeMetric("testxx")

	return conf, nil
}

func (p *parser) Merge(parent interface{}, child interface{}) interface{} {
	panic("TODO")
}

func configFactory(c interface{}) api.StreamFilterFactory {
	conf, ok := c.(*config)
	if !ok {
		panic("unexpected config type")
	}
	return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &filter{
			callbacks: callbacks,
			config:    conf,
		}
	}
}
