package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()       //means let the CPU execute other goroutines, and come back at some point.
        fmt.Println(s)
    }
}

func main() {
    go say("world") // create a new goroutine
    say("hello") // current goroutine
}