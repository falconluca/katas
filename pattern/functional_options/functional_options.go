package functional_options

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServerFo(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func Fo() {
	s1, _ := NewServerFo("localhost", 1024)
	fmt.Println(s1)

	s2, _ := NewServerFo("localhost", 2048, Protocol("udp"))
	fmt.Println(s2)

	s3, _ := NewServerFo("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
	fmt.Println(s3)
}
