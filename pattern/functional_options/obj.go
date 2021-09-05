package functional_options

import (
	"crypto/tls"
	"fmt"
	"time"
)

type ServerObj struct {
	Addr string
	Port int
	Conf *ServerConfig
}

type ServerConfig struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

func NewServer(addr string, port int, conf *ServerConfig) (*ServerObj, error) {
	if conf == nil {
		defaultConf := ServerConfig{Protocol: "tcp"}
		return &ServerObj{Addr: addr, Port: port, Conf: &defaultConf}, nil
	}
	return &ServerObj{Addr: addr, Port: port, Conf: conf}, nil
}

func Obj() {
	srv1, _ := NewServer("localhost", 9000, nil)
	fmt.Println(srv1) // TODO toString

	conf := ServerConfig{Protocol: "tcp"}
	srv2, _ := NewServer("localhost", 9000, &conf)
	fmt.Println(srv2)
}
