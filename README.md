# Go-Mask - simple Go module that provides functionality to mask strings

[![GoDoc](https://godoc.org/github.com/vlaship/go-mask?status.svg)](https://godoc.org/github.com/vlaship/go-mask)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlaship/go-mask)](https://goreportcard.com/report/github.com/vlaship/go-mask)
[![Build Status](https://github.com/vlaship/go-mask/workflows/build/badge.svg)](https://github.com/vlaship/go-mask/actions)
[![Coverage Status](https://coveralls.io/repos/github/vlaship/go-mask/badge.svg?branch=main)](https://coveralls.io/github/vlaship/go-mask?branch=main)

## Installation

To install Go-Mask, you can use `go get`:

```bash
go get github.com/vlaship/go-mask# go-mask
```

## Usage

Import the package in your Go file:

```go
import "github.com/vlaship/go-mask"
```

Use it like this:

```go
masked := mask.String("1234567890")
```

