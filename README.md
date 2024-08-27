# DEPRECATED

The full node reference implementation was [rewritten in Rust](https://github.com/kaspanet/rusty-kaspa), as a result, the Go implementation is now deprecated.

PLEASE NOTE: Any pull requests or issues that will be opened in this repository will be closed without treatment, except for issues or pull requests related to the bricswallet, which remains maintained. In any other case, please use the [Rust implementation](https://github.com/kaspanet/rusty-kaspa) instead.

# Bricsd

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/brics18/bricsd)

Bricsd was the reference full node Brics implementation written in Go (golang).

## What is Brics

Brics is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install bricsd including all dependencies:

```bash
$ git clone https://github.com/brics18/bricsd
$ cd bricsd
$ go install . ./cmd/...
```

- Bricsd (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

## Getting Started

Bricsd has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ bricsd
```

## Discord

Join our discord server using the following link: https://discord.gg/YNYnNN5Pf2

## Issue Tracker

The [integrated github issue tracker](https://github.com/brics18/bricsd/issues)
is used for this project.

Issue priorities may be seen at https://github.com/orgs/kaspanet/projects/4

## Documentation

The [documentation](https://github.com/kaspanet/docs) is a work-in-progress

## License

Bricsd is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
