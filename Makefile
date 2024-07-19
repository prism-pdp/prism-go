IMAGE_NAME = dpduado-go

.PHONY: test

build-img:
	docker build -t $(IMAGE_NAME) .

shell:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build --entrypoint bash $(IMAGE_NAME)

test:
	docker run -it --rm -v .:/app -v /tmp/go-cache:/root/.cache/go-build $(IMAGE_NAME) test ./xz21
