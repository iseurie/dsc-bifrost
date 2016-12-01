package main

import (
    "os"
    "fmt"
    "path"
    "flag"
    "bitbucket.org/iseurie/bifrost/bridge"
)

func main() {
    var cfgPath string
    flag.Parse()
    switch(len(flag.Args())) {
    case 1: cfgPath = flag.Args()[0]; break
    default: cfgPath = "./config.json"
    }

    err := bridge.MainBrIndex.IfParse(cfgPath)
    if(err != nil) {
        samplePath := path.Clean(cfgPath) + "/config.json"
        cfi, nerr := os.Stat(cfgPath)
        if nerr != nil { 
            fmt.Fprintf(os.Stderr, "%s%v\n", 
                    "Cannot index bridge configuration: ", err)
            os.Exit(1)
        } else if cfi.IsDir() {
            fmt.Printf("%s%s%s%c\n", "Stat found directory; ",
                    "generating sample configuration at '",
                    samplePath, '\'')
            err = bridge.WriteSample(samplePath)
            if err != nil {
                fmt.Fprintf(os.Stderr, "%s%v\n", 
                        "Cannot generate sample configuration: ", err)
                os.Exit(1)
            }; return
        }
    }
    fmt.Println("Bridge configuration read. Direct coordinates:")
    for _, dbr := range bridge.MainBrIndex.DirectBrs {
        fmt.Printf("\t%s <=> %s@[", dbr.DscC.Email,
                dbr.IrcC.CtrlNick)
        for i, h := range dbr.IrcC.Hosts {
            fmt.Print(h.Host)
            if i < len(dbr.IrcC.Hosts)-1 { fmt.Print(", ") }
        }; fmt.Print("] : [")
        for i, w := range dbr.IrcC.Wranglers {
            fmt.Print(w)
            if i < len(dbr.IrcC.Wranglers)-1 { fmt.Print(", ") }
        }; fmt.Println("]")
    }
    fmt.Println("Setting up...")
    initBrs(&bridge.MainBrIndex)
    fmt.Println("Now online.")
}
