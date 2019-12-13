clpwrapper
==========

This is a small package that provides interoperability
among:
* [clp](https://github.com/lanl/clp) for solving linear programs
* [Gonum mat](https://github.com/gonum/gonum/mat) for
general matrices
* [sparse](https://github.com/james-bowman/sparse) for sparse matrices

clp's "natural" matrix type comes from [CoinUtils](https://github.com/coin-or/CoinUtils) and is fine in that limited
context. But it is hardly ideal for interacting with the
larger go ecosystem.

## Install

clpwrapper requires all three packages listed above, including
their dependencies (clp's in particular are non-trivial).
Then run:

```bash
go get github.com/jonreiter/clpwrapper
```

## Documentation

Docs are in the [usual place](https://godoc.org/github.com/jonreiter/clpwrapper) thanks to the fantastic work of the
[GoDoc](https://godoc.org/) project.
