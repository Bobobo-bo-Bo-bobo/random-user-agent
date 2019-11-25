package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Note: I'm are aware of the fact that this does not generate "safe" randoms.
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println(UserAgents[rand.Intn(len(UserAgents))])
}

