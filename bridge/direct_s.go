package bridge

import (
    "github.com/bwmarrin/discordgo"
    irccl "github.com/fluffle/goirc/client"
)

type IrcCoordinate struct {
    Hosts []NetHost         `json:"hosts"`                  //networks on which to init
    Proxy NetHost           `json:"proxy,omitempty"`        //optional
    Wranglers []string      `json:"wranglers"`              //nicks authorized for this bridge
    CtrlNick string         `json:"nick"`                   //nick for controller interface
    CtrlPass string         `json:"pass"`                   //use where nick non-unique
}

type DscCoordinate struct {
    Email string    `json:"email"`
    Pass string     `json:"pass"`
}

type DirectBridge struct {
    IrcC IrcCoordinate          `json:"irc"`
    DscC DscCoordinate          `json:"dsc"`
    Session *discordgo.Session  `json:"-"`
    Spoofs map[string]ircNode   `json:"-"`  //key by ident/Discord UID
    Control ircNode             `json:"-"`
    Up chan bool                `json:"-"`
}

type ircNode struct {
    Target *discordgo.User
    EvtRmvers map[string]*irccl.Remover     //key by event name
    Clients map[string]*irccl.Conn          //key by hostname
}
