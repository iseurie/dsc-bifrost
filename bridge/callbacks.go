package bridge

import (
    "github.com/bwmarrin/discordgo"
    irccl "github.com/fluffle/goirc/client"
)

func onDiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
    for _, dbr := range MainBrIndex.DirectBrs {
        target, _ := s.User("@me")
        if sp, ok := dbr.Spoofs[target.ID]; ok {
             for _, w := range dbr.IrcC.Wranglers {
                for _, c := range sp.Clients {
                    c.Privmsg(w, m.ContentWithMentionsReplaced())
                }
            }
        }
    }; return
}        

//IRC spoof message callback
func onSpoofMessage(c *irccl.Conn, l *irccl.Line) {
    if l.Public() { return }
    for _, dbr := range MainBrIndex.DirectBrs {
        for _, node := range dbr.Spoofs {
            chls, _ := dbr.Session.UserChannels()
            for _, ch := range chls {
                if ch.Recipient == node.Target {
                    dbr.Session.ChannelMessageSend(ch.ID, l.Text())
                }
            }
        }
    }; return
}
