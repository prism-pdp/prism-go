IMAGE_NAME = prism-go

.PHONY: test

build-img:
	docker compose build go

shell:
	$(MAKE) run OPT="--entrypoint bash"

run:
	docker compose run --rm $(OPT) go $(CMD)

test/clean:
	$(MAKE) run CMD="clean -testcache"

test:
	$(MAKE) test/clean
	dd if=/dev/zero bs=1M count=100 > xz21/testdata/dummy.data 2> /dev/null
	$(MAKE) run CMD="test -v ./xz21"

