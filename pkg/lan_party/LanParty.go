package lanparty

import (
	"lanyard/pkg/lan_party/gamelist"
	"time"
)

type LanParty struct {
	Options         map[string]any
	DepartureTimes  []time.Time
	Usernames       map[int64]string
	Identities      map[string]string
	UserResolutions map[int64]gamelist.GameListResolution
}
