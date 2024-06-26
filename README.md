# dpduado-go

## Description

dpduado-go is a Go project for dpduado.

## Tips for developers

### Go bindings

This project uses Go bindings created in [dpduado-sol][link:dpduado-sol].
If you update solidity codes, you should copy the Go bindings of updated codes to `sol` directory.

### Testing

Test codes should be placed in `sol` directory with `_test` suffixed to the filename.
You can run test with following command.

```sh
make test
```

[link:dpduado-sol]: https://github.com/dpduado/dpduado-sol "dpduado-sol"