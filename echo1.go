package main


import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
            fmt.Println("num->"+os.Args[i])
            s +=sep + os.Args[i]
            sep = "_"
        }
        fmt.Println("init:"+os.Args[0])
        fmt.Println("Total:"+s)
        }
