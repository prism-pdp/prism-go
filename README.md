# prism-go

*prism-go* is a cryptographic library for PRISM, which is a Provable Data Possession (PDP) system.

## Features

- Supports data verification based on PDP
- Lightweight design using Go
- Easy setup with Docker
- Cryptographic support via [the PBC library](https://crypto.stanford.edu/pbc/)

## Technologies Used

- [**Go**](https://go.dev/): For main programming language
- [**PBC library**](https://crypto.stanford.edu/pbc/): For cryptographic library
- [**PBC Go wrapper**](https://pkg.go.dev/github.com/nik-u/pbc): For Go wrapper of the PBC library 
- [**Docker**](https://www.docker.com/): For environment virtualization and management

## System Requirements

This project has been tested and is verified to work on **Ubuntu 24.04.01 LTS**.
We recommend using this version of Ubuntu for the best experience. 
Other operating systems or versions might require additional setup or may not be supported.

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

