package main

import (
    "fmt"

    "github.com/i0Ek3/color"
    log "github.com/sirupsen/logrus"
)

func main() {
    fmt.Println(color.Blue("This is a Git Flow test."))
    fmt.Println(color.Red("Please don't use it!"))
    fmt.Println(color.Gray("What a taugh day!"))
    fmt.Println(color.Yellow("This message created by @deadl10n."))
    log.Warnf("Add log by @i0Ek3.")
    fmt.Println(color.Colorize("cyan", "add new lib"))
}
