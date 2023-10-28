[![GoDoc](https://img.shields.io/badge/go.dev-reference-blue)](https://pkg.go.dev/github.com/chmike/domain)
![Build](https://github.com/chmike/domain/actions/workflows/audit.yml/badge.svg)
[![codecov](https://codecov.io/gh/chmike/domain/graph/badge.svg?token=06TJPZ1S5J)](https://codecov.io/gh/chmike/domain)
[![Go Report](https://goreportcard.com/badge/github.com/chmike/domain)](https://goreportcard.com/report/github.com/chmike/domain)
![Status](https://img.shields.io/badge/status-stable-brightgreen.svg)
![release](https://img.shields.io/github/release/chmike/domain.svg)

# domain.Check()

This package contains a single function that checks the validity of a domain name. 

A domain name must respect rules defined in 
- [section 3.5 of RFC 1034 ("Domain names - concepts and facilities")](https://tools.ietf.org/html/rfc1034#section-3.5)
- [section 2 of RFC 1123 ("Requirements for Internet Hosts -- Application and Support")](https://tools.ietf.org/html/rfc1123#section-2)

The `domain.Check` function ensures that the domain name respect those rules. If not, it returns an error explaining the detected problem. 

## Prerequisites

The package has no prerequisites and external dependencies.

## Installation

To install or update this package use the instruction:

```bash
go get -u "github.com/chmike/domain"
```

## Usage examples

The `Check` function can be used to check the validity of host or domain names.

```go
name := "host.example.com"
if err := domain.Check(name); if err != nil {
    log.Fatalf("invalid domain name '%s': %v", name, err)
}
```


