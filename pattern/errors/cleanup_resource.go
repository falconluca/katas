package errors

import (
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CleanupResource() {
	file := "/Users/luca/go/src/katas/hello.go"
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("error opening '%s', err: %v\n", file, err)
	}
	// 使用 defer 关键词进行清理
	defer Close(f)

	file = "/Users/luca/go/src/katas/if.go"
	f, err = os.Open(file)
	if err != nil {
		log.Fatalf("error opening '%s', err: %v\n", file, err)
	}
	defer Close(f)
}
