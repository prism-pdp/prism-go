# prism-go

*prism-go* is a cryptographic library for Provable Data Possession (PDP).

## Features

- Supports data verification based on PDP
- Lightweight design using Go
- Easy setup with Docker
- Cryptographic support via [the PBC library](https://crypto.stanford.edu/pbc/)

## Technologies Used

- [**Go**](https://go.dev/): Main programming language
- [**PBC library**](https://crypto.stanford.edu/pbc/): Cryptographic library
- [**PBC Go wrapper**](https://pkg.go.dev/github.com/nik-u/pbc): Go wrapper of the PBC library 
- [**Docker**](https://www.docker.com/): Virtualization environement

## Installation

### 1. Install Docker

Follow the [official guide](https://docs.docker.com/get-docker/) to install Docker.

### 2. Clone the repository

```bash
git clone https://github.com/prism-pdp/prism-go.git
```

### 3. Build the docker image

```bash
make build-img
```

## Testing

You can run test with following command.

```sh
make test
```

