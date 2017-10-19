# random
random utils for go
## random string 
generates an alpha numeric random string 
#### example
    ```go
        package main

        import (
            "random"
            "fmt"
        )
        
        func main() {
            fmt.Println(random.String(10)) //tq9Af6hv5T
        }
    ```