<p align="center">
  <h1 align="center">ERRNO</h1>
  <p align="center">C-like err codes providable module</p>
</p>

<div align="center">

  [![Test & Build](https://github.com/rejchev/errno/actions/workflows/build.yml/badge.svg)](https://github.com/rejchev/errno/actions/workflows/build.yml)
  [![Go Reference](https://pkg.go.dev/badge/github.com/rejchev/errno.svg)](https://pkg.go.dev/github.com/rejchev/errno)
  [![Go Report Card](https://goreportcard.com/badge/github.com/rejchev/errno)](https://goreportcard.com/report/github.com/rejchev/errno)
  
</div>

### About
C-like err codes providable module

### Installation
```bash
go get github.com/rejchev/errno
```

### Usage example
```go
import (
	"context"

	errnov1 "github.com/rejchev/errno"
)

func main() {

	buff := []byte{} 
	if code := read(&buff); FAIL(code) {
		panic(code.String())
	}

}

func read(buff *[]byte) errnov1.Code {
	if buff == nil {
		return errnov1.EINVAL
	}

	// do something ...

	return errnov1.OK
}
```