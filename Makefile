IMAGE_NAME = dpduado-go

.PHONY: test

build-img:
	docker build -t $(IMAGE_NAME) .

shell:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build --entrypoint bash $(IMAGE_NAME)

test/clean:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) clean -testcache

test:
	$(MAKE) test/clean
	dd if=/dev/zero bs=1M count=100 > xz21/testdata/dummy.data 2> /dev/null
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test -v ./xz21

test1:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) clean -testcache
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test1

test2:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test2

test3:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21 -run Test3
