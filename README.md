# go-protected

[![tag](https://img.shields.io/github/tag/peczenyj/go-protected.svg)](https://github.com/peczenyj/go-protected/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://pkg.go.dev/badge/github.com/peczenyj/go-protected)](http://pkg.go.dev/github.com/peczenyj/go-protected)
[![Go](https://github.com/peczenyj/go-protected/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/go-protected/actions/workflows/go.yml)
[![Lint](https://github.com/peczenyj/go-protected/actions/workflows/lint.yml/badge.svg)](https://github.com/peczenyj/go-protected/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/peczenyj/go-protected/graph/badge.svg?token=9y6f3vGgpr)](https://codecov.io/gh/peczenyj/go-protected)
[![Report card](https://goreportcard.com/badge/github.com/peczenyj/go-protected)](https://goreportcard.com/report/github.com/peczenyj/go-protected)
[![CodeQL](https://github.com/peczenyj/go-protected/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/peczenyj/go-protected/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/peczenyj/go-protected/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/peczenyj/go-protected/actions/workflows/dependency-review.yml)
[![License](https://img.shields.io/github/license/peczenyj/go-protected)](./LICENSE)
[![Latest release](https://img.shields.io/github/release/peczenyj/go-protected.svg)](https://github.com/peczenyj/go-protected/releases/latest)
[![GitHub Release Date](https://img.shields.io/github/release-date/peczenyj/go-protected.svg)](https://github.com/peczenyj/go-protected/releases/latest)
[![Last commit](https://img.shields.io/github/last-commit/peczenyj/go-protected.svg)](https://github.com/peczenyj/go-protected/commit/HEAD)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/peczenyj/go-protected/blob/main/CONTRIBUTING.md#pull-request-process)

yet another mutex-protected variable in golang

## usage

```go
    lockVar := goprotected.New(new(int)) // import "github.com/peczenyj/go-protected"

    lockVar.Use(func(i *int) {
        *i++
    })

    lockVar.Use(func(i *int) {
        fmt.Println("got i=", *i)
    })

    // Output:
    // got i= 1
```
