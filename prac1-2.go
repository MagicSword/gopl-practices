package main


import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
            fmt.Sprintf(" %d -> %s ",i,os.Args[i])
            s +=sep + os.Args[i]
            sep = " "
        }
        fmt.Println("os.Args[0] = " + os.Args[0])
        fmt.Println("Total = " + s)
        }
