package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
    "fmt"
)

func WordCount(s string) map[string]int {
    fmt.Println(s)
    var count map[string]int = make(map[string]int)
    fields := strings.Fields(s)
    fmt.Printf("%q", fields)
    for _, x := range fields {
        fmt.Println(x)
        count[x]++
    }
    
    return count
}

func main() {
    wc.Test(WordCount)
}

