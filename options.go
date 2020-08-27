package jsonrpc

import (
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

type ParamEncoder func(reflect.Value) (reflect.Value, error)

type Config struct {
	reconnectBackoff backoff

	paramEncoders map[reflect.Type]ParamEncoder

	proxyConnFactory func(func() (*websocket.Conn, error)) func() (*websocket.Conn, error) // for testing
}

func defaultConfig() Config {
	return Config{
		reconnectBackoff: backoff{
			minDelay: 100 * time.Millisecond,
			maxDelay: 5 * time.Second,
		},

		paramEncoders: map[reflect.Type]ParamEncoder{},
	}
}

type Option func(c *Config)

func WithReconnectBackoff(minDelay, maxDelay time.Duration) func(c *Config) {
	return func(c *Config) {
		c.reconnectBackoff = backoff{
			minDelay: minDelay,
			maxDelay: maxDelay,
		}
	}
}

func WithParamEncoder(t interface{}, encoder ParamEncoder) func(c *Config) {
	return func(c *Config) {
		c.paramEncoders[reflect.TypeOf(t).Elem()] = encoder
	}
}
