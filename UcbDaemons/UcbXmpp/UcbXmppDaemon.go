package UcbXmpp

import (
	"github.com/alfredyang1986/blackmirror/bmxmpp"
	"os"
)

type UcbXmpp struct {
	//GroupRoomID string
	XmppConfig bmxmpp.BmXmppConfig
}

func (r UcbXmpp) NewUcbXmppBDaemon(args map[string]string) *UcbXmpp {
	env := os.Getenv("BM_XMPP_CONF_HOME") + "/resource/xmppconfig.json"
	os.Setenv("BM_XMPP_CONF_HOME", env)
	bxc, _ := bmxmpp.GetConfigInstance()
	ins := UcbXmpp {
		//GroupRoomID: args["room"],
		XmppConfig: *bxc,
	}
	return &ins
}

func (r UcbXmpp) SendGroupMsg(room, msg string) error {
	return r.XmppConfig.Forward2Group(room, msg)
}


