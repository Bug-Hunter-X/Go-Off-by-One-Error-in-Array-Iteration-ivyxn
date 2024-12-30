```go
package main

import "fmt"

func main() {
    var a [5]int
    for i := 0; i <= 5; i++ {
        a[i] = i * 2
    }
    fmt.Println(a)
}
```