IMAGE_NAME = prism-go

.PHONY: test

build-img:
	docker compose build go

shell:
	docker compose run --entrypoint bash go

test/clean:
	docker compose run go clean -testcache

test:
	$(MAKE) test/clean
	dd if=/dev/zero bs=1M count=100 > xz21/testdata/dummy.data 2> /dev/null
	docker compose run go test -v ./xz21

test1:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) clean -testcache
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test1

test2:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test2

test3:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test3
