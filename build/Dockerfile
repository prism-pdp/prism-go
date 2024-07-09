FROM golang:1.22.2-bookworm

RUN apt update \
	&& apt -y upgrade

RUN apt -y install \
	vim

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go mod download

COPY --chmod=755 ./assets/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

