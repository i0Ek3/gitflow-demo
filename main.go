package main

import (
    "fmt"

    "github.com/i0Ek3/color"
    log "github.com/sirupsen/logrus"
)

func main() {
    fmt.Println(color.Blue("This is a Git Flow test."))
    fmt.Println(color.Red("Please don't use it!"))
    fmt.Println(color.White("What a taugh day!"))
    fmt.Println(color.Yellow("This message created by @deadl10n."))
    log.Warnf("Add log by @i0Ek3.")
    fmt.Println(color.Colorize("cyan", "add new lib"))
    hi("deadl10n")
    debug(hi)
}

func hi(name string) {
    fmt.Println("Hi,", name)
} 

func debug(msg ...any) {
    s := fmt.Sprintf("------> %s::", msg)
    fmt.Println(s)
}
