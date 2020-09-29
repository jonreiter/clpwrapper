# clpwrapper

[![GoDoc](https://godoc.org/github.com/jonreiter/clpwrapper?status.svg)](https://godoc.org/github.com/jonreiter/clpwrapper)
[![Go Report](https://goreportcard.com/badge/github.com/jonreiter/clpwrapper)](https://goreportcard.com/report/github.com/jonreiter/clpwrapper)
[![BuildStatus](https://www.travis-ci.com/jonreiter/clpwrapper.svg?branch=master)](https://www.travis-ci.com/github/jonreiter/clpwrapper/branches)
[![codecov](https://codecov.io/gh/jonreiter/clpwrapper/branch/master/graph/badge.svg)](https://codecov.io/gh/jonreiter/clpwrapper)


This is a small package that provides interoperability
among:

* [clp](https://github.com/lanl/clp) for solving linear programs
* [Gonum mat](https://github.com/gonum/gonum/) for
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
