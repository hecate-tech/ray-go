.PHONY: all

all:
	c-for-go ray.yml

clean:
	rm -f ray/cgo_helpers.c ray/cgo_helpers.h ray/cgo_helpers.go
	rm -f ray/const.go
	rm -f ray/doc.go
	rm -f ray/ray.go
	rm -f ray/types.go

test:
	go build