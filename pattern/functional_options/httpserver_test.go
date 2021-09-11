package functional_options

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHttpServer(t *testing.T) {
	assert := assert.New(t)

	server, _ := NewDefaultHttpServer("127.0.0.1", 80)
	assert.Equal("127.0.0.1", server.Addr)
	assert.Equal(80, server.Port)
	assert.Equal("tcp", server.Protocol)
	assert.Nil(server.TLS)
}

func TestNewTLSHttpServer(t *testing.T) {
	assert := assert.New(t)

	conf := &tls.Config{}
	server, _ := NewTLSHttpServer("127.0.0.1", 80, conf)
	assert.Equal("127.0.0.1", server.Addr)
	assert.Equal(80, server.Port)
	assert.Equal("tcp", server.Protocol)
	assert.NotNil(server.TLS)
}

func TestNewHttpServerWithTimeout(t *testing.T) {
	assert := assert.New(t)

	timeout := 30 * time.Second
	server, _ := NewHttpServerWithTimeout("127.0.0.1", 80, timeout)
	assert.Equal(timeout, server.Timeout)
}

func TestNewTLSHttpServerWithMaxConnAndTimeout(t *testing.T) {
	assert := assert.New(t)

	timeout := 30 * time.Second
	conf := &tls.Config{}
	server, _ := NewTLSHttpServerWithMaxConnAndTimeout("127.0.0.1", 80, 100, timeout, conf)
	assert.Equal(timeout, server.Timeout)
	assert.NotNil(server.TLS)
}
