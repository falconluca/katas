# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# .PHONY标记声明不检查文件名
.PHONY: build clean help

all: build

build:
	go build -v .

clean:
	rm -rf blogapi
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"