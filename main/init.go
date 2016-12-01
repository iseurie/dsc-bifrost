package main

import (
    "fmt"
    "os"
    "bitbucket.org/iseurie/bifrost/bridge"
)

func initBrs(brs *bridge.BrIndex) {
    for _, br := range brs.DirectBrs {
        fmt.Printf("Authenticating with Discord as '%s'...\n", br.DscC.Email)
        err := br.Init()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Initialization failed: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Scaffolding control '%s':\n", br.IrcC.CtrlNick)
        for _, h := range br.IrcC.Hosts {
            fmt.Printf("\t@%s:%d\n...", h.Host, h.Port)
            err := br.MapControl(&h)
            if err != nil {
                fmt.Fprintf(os.Stderr, "\tFailed mapping to '%s': %v\n", h.String(), err)
            }
        }
    }
}
