FROM golang:1.23.2-bookworm

RUN apt update \
	&& apt -y upgrade

RUN apt -y install \
	libgmp-dev \
	build-essential \
	flex \
	bison

WORKDIR /usr/src/pbc-0.5.14
RUN curl -sSL https://crypto.stanford.edu/pbc/files/pbc-0.5.14.tar.gz | tar zx --strip-components 1 -C . \
	&& ./configure \
	&& make \
	&& make install

RUN apt -y install \
	vim

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go mod download

ENTRYPOINT ["go"]
