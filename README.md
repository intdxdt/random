# random
random utils for go
## random string 
generates an alpha numeric random string 
```go
    package main

    import (
        "fmt"
        "github.com/intdxdt/random"
    )
    
    func main() {
        fmt.Println(random.String(10)) //tq9Af6hv5T
    }
```