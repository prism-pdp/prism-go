IMAGE_NAME = dpduado-go

.PHONY: test

build-img:
	docker build -t $(IMAGE_NAME) .

shell:
	docker run -it --rm --volume .:/app --entrypoint bash $(IMAGE_NAME)

test:
	docker run -it --rm --volume .:/app $(IMAGE_NAME) test
