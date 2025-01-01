# prism-go

## Description

prism-go is a Go project for prism.

## Tips for developers

### Go bindings

This project uses Go bindings created in [prism-sol][link:prism-sol].
If you update solidity codes, you should copy the Go bindings of updated codes to `sol` directory.

### Testing

Test codes should be placed in `sol` directory with `_test` suffixed to the filename.
You can run test with following command.

```sh
make test
```

[link:prism-sol]: https://github.com/prism/prism-sol "prism-sol"