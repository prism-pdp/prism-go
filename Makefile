.PHONY: test

docker-build:
	docker compose build go-compiler

docker-run:
	docker compose run -it --rm go-compiler $(CMD)

test:
	make docker-run CMD='go test -v ./sol'

shell:
	make docker-run CMD='/bin/bash'
