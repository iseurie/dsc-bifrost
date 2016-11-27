package bridge

import (
    "encoding/json"
    "io/ioutil"
)

var MainBrIndex BrIndex

// Represents a predefined bridge configuration
type BrIndex struct {
    DirectBrs []DirectBridge    `json:"direct"`
}

// Represents an IRC host as specified in runtime configuration
type NetHost struct {
    Host string         `json:"host"`
    Pass string         `json:"pass"`
    Port uint16         `json:"port"`
}

func(h *NetHost) String() string {
    return h.Host + ":" + string(h.Port)
}

// Read the configuration at 'path' into the given struct
func(bri *BrIndex) IfParse(path string) (err error) {
    dat, err := ioutil.ReadFile(path)
    if err != nil { return err }
    err = json.Unmarshal(dat, bri); return
}

// Generate sample configuration
func WriteSample(path string) (err error) {
    sampleBr := BrIndex {
        DirectBrs : []DirectBridge{
            {
                IrcC : IrcCoordinate{
                    Hosts : []NetHost{
                        { 
                            Host : "irc.freenode.net",
                            Pass : "",
                            Port : 6667,
                        },
                    },

                    Wranglers : []string{ "your nick", "your alt's" },
                    CtrlNick : "houston",
                    CtrlPass : "secure_psk",
                },
                DscC : DscCoordinate{ "you@example.com", "discord_pass" },
            },
        },
    }
    dat, err := json.MarshalIndent(sampleBr, "", "    ")
    if err != nil { return err }
    err = ioutil.WriteFile(path, dat, 0644); return
}
