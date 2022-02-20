go-windows-junction
===================

Create junctions on Windows.

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

```
$ go run example.go ../nyagos nya
nya  is linked to  C:\Users\hymko\go\src\github.com\nyaosorg\nyagos
```

