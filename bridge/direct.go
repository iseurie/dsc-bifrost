package bridge

import (
    "github.com/bwmarrin/discordgo"
    irccl "github.com/fluffle/goirc/client"
    ircst "github.com/fluffle/goirc/state"
)

func(dbr *DirectBridge) Init() (err error) {
    dbr.Session, err = discordgo.New(dbr.DscC.Email, dbr.DscC.Pass)
    dbr.Spoofs = make(map[string]ircNode)
    dbr.Control.Clients = make(map[string]*irccl.Conn)
    return
}

// Add an IRC controller and connect spoofs by host
func(dbr *DirectBridge) MapControl(h *NetHost) (err error) {
    dscTarget, err := dbr.Session.User("@me")
    if err != nil { return err }
    mNick := new(ircst.Nick)
    mNick.Nick = dbr.IrcC.CtrlNick
    mNick.Ident = dscTarget.ID
    
    nCfg := irccl.NewConfig(dbr.IrcC.CtrlNick)
    nCfg.Me = mNick
    nCfg.Server = h.Host + ":" + string(h.Port)
    nCfg.Pass = h.Pass

    ctrl := irccl.Client(nCfg)
    if err != nil { return err }
    dbr.Control.Clients[h.Host] = ctrl
    dbr.Control.Target = dscTarget
    for _, s := range dbr.Spoofs {
        err = s.MapSpoof(h)
        if err != nil { return err }
    }; return
}

// Initiate IRC control; add event hooks
func(dbr *DirectBridge) SetUp() (err error) {
    dbr.Session, err = discordgo.New(dbr.DscC.Email, dbr.DscC.Pass)
    if err != nil { return err }
    err = dbr.Open()
    if err != nil { return err }
    // init discord callbacks
    dbr.Session.AddHandler(onDiscordMessage)
    for _, h := range dbr.IrcC.Hosts {
        dscTarget, err := dbr.Session.User("@me")
        if err != nil { return err }
        dbr.Control.Init(dscTarget)
        err = dbr.Control.MapSpoof(&h)
        for _, cl := range dbr.Control.Clients {
            err = cl.Connect()
            if err != nil { return err }
        }
    }; return
}

// Tear down all IRC and Discord clients
func(dbr *DirectBridge) SetDown() (err error) {
    for _, s := range dbr.Spoofs {
        for _, cl := range s.Clients {
            cl.Quit()
        }
    }; for _, cl := range dbr.Control.Clients {
        cl.Quit()
    }; err = dbr.Session.Close(); return
}
