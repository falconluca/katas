package functional_options

import (
	"crypto/tls"
	"time"
)

// Go不支持重载函数, 所以得用不同的函数名来应对不同的配置选项
type HttpServer struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

func NewDefaultHttpServer(addr string, port int) (*HttpServer, error) {
	httpServer := HttpServer{
		addr,
		port,
		"tcp",
		30 * time.Second,
		100,
		nil,
	}
	return &httpServer, nil
}

func NewTLSHttpServer(addr string, port int, tls *tls.Config) (*HttpServer, error) {
	return &HttpServer{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

func NewHttpServerWithTimeout(addr string, port int, timeout time.Duration) (*HttpServer, error) {
	return &HttpServer{addr, port, "tcp", timeout, 100, nil}, nil
}

func NewTLSHttpServerWithMaxConnAndTimeout(addr string, port int,
	maxconns int, timeout time.Duration, tls *tls.Config) (*HttpServer, error) {
	return &HttpServer{addr, port, "tcp", timeout, maxconns, tls}, nil
}
