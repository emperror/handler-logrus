# Logrus handler

[![CircleCI](https://circleci.com/gh/emperror/handler-logrus.svg?style=svg)](https://circleci.com/gh/emperror/handler-logrus)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/logrus?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/logrus)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-logrus.svg)](https://golangci.com/r/github.com/emperror/handler-logrus)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/emperror.dev/handler/logrus)

**Error handler integration for [Logrus](https://github.com/sirupsen/logrus).**


## Installation

```bash
go get emperror.dev/handler/logrus
```


## Usage

```go
package main

import (
	"github.com/sirupsen/logrus"

	logrushandler "emperror.dev/handler/logrus"
)

func main() {
	logger := logrus.New()
	handler := logrushandler.New(logger)
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
