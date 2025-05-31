go-windows-junction
===================

[![Go Reference](https://pkg.go.dev/badge/github.com/nyaosorg/go-windows-junction.svg)](https://pkg.go.dev/github.com/nyaosorg/go-windows-junction)
[![Go Test](https://github.com/nyaosorg/go-windows-junction/actions/workflows/go.yml/badge.svg)](https://github.com/nyaosorg/go-windows-junction/actions/workflows/go.yml)

A Go package to create [junction points](https://docs.microsoft.com/en-us/windows/win32/fileio/hard-links-and-junctions) on Windows.

- On **Windows**, it creates a junction point using native system calls.
- On **non-Windows** platforms (Linux, macOS), it falls back to `os.Symlink`.

Junctions are similar to symbolic links but have certain limitations:
- They can only point to directories.
- They do not require elevated privileges (unlike symlinks on some versions of Windows).
- The link target can be read using `os.Readlink`.

This package is useful when writing cross-platform code that creates directory links, and you want consistent behavior on Windows without requiring admin rights.

## Installation

```sh
go get github.com/nyaosorg/go-windows-junction
````

## Usage Example

```go
package main

import (
    "os"
    "errors"

    "github.com/nyaosorg/go-windows-junction"
)

func main1() error {
    if len(os.Args) < 3 {
        return errors.New("go run example.go TARGET MOUNTPT")
    }
    if err := junction.Create(os.Args[1], os.Args[2]); err != nil {
        return err
    }
    result, err := os.Readlink(os.Args[2])
    if err != nil {
        return err
    }
    println(os.Args[2], " is linked to ", result)
    return nil
}

func main() {
    if err := main1(); err != nil {
        println(err.Error())
        os.Exit(1)
    }
}
```

### Output Example

```
$ go run example.go ../nyagos nya
nya  is linked to  C:\Users\hymko\go\src\github.com\nyaosorg\nyagos
```

## Author

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

## License

MIT License
