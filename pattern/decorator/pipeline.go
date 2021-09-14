package decorator

import "net/http"

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

// Handler 工具函数用来遍历并调用各个decorator pipeline
func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}
