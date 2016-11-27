package bridge

import (
    "github.com/bwmarrin/discordgo"
    irccl "github.com/fluffle/goirc/client"
    "regexp"
)

//parse a valid IRC nick from a given discriminator
func ParseSpoofKey(disc string, exp *string) string {
    var allowed *regexp.Regexp
    if exp == nil {
        tags := regexp.MustCompile("[#]")
        disc = tags.ReplaceAllString(disc, "_");
        allowed = regexp.MustCompile("[^\\p{IsAlphabetic}^\\p{IsDigit}^_]")
    } else {
        allowed = regexp.MustCompile(*exp)
    }; return allowed.FindString(disc)
}

//Initialize an IRC endpoint struct
func(n *ircNode) Init(target *discordgo.User) {
    n.Target = target
    n.Clients = make(map[string]*irccl.Conn)
    n.EvtRmvers = make(map[string]*irccl.Remover)
}

//Connect and map an IRC endpoint to a spoof on a given host
func(s *ircNode) MapSpoof(h *NetHost) (err error) {
    var spoof_evts = [...]string{ irccl.PRIVMSG, irccl.ACTION, }
    ncfg := irccl.NewConfig(ParseSpoofKey(s.Target.Discriminator, nil))
    ncfg.Me.Ident = s.Target.ID
    ncfg.Server = h.Host + ":" + string(h.Port)
    ncfg.Pass  = h.Pass
    s.Clients[h.Host] = irccl.Client(ncfg)
    if err != nil { return err }
    for _, e := range spoof_evts {
        rmver := s.Clients[h.Host].HandleFunc(e, onSpoofMessage)
        s.EvtRmvers[e] = &rmver
    }; return
}
