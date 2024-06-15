package commands

import (
	"context"
	"fmt"
	cmds "github.com/bittorrent/go-btfs-cmds"
	"github.com/bittorrent/go-btfs/repo/fsrepo"
	"github.com/libp2p/go-libp2p/p2p/host/peerstore/pstoreds"
)

// ErrDepthLimitExceeded indicates that the max depth has been exceeded.

var PeerStore = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:          "",
		ShortDescription: ``,
		LongDescription:  ``,
	},

	Subcommands: map[string]*cmds.Command{
		"ls":     lsCmd,
		"delete": deleteCmd,
		"clean":  cleanCmd,
	},
}

var lsCmd = &cmds.Command{
	Arguments: []cmds.Argument{
		{
			Name: "path",
		},
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		repo, _ := fsrepo.Open(request.Arguments[0])
		pstore, _ := pstoreds.NewPeerstore(context.Background(), repo.Datastore(), pstoreds.DefaultOpts())
		address := pstore.PeersWithAddrs()
		for _, x := range address {
			fmt.Println(x.String())
		}
		return nil
	},
}

var deleteCmd = &cmds.Command{
	Arguments: []cmds.Argument{
		{
			Name: "path",
		},
		{
			Name: "peerId",
		},
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		repo, _ := fsrepo.Open(request.Arguments[0])
		pstore, _ := pstoreds.NewPeerstore(context.Background(), repo.Datastore(), pstoreds.DefaultOpts())
		// 获取参数值
		peerId := request.Arguments[1]
		peers := pstore.PeersWithAddrs()
		for _, p := range peers {
			if p.String() == peerId {
				pstore.ClearAddrs(p)
				break
			}
		}
		fmt.Println(peerId, "was deleted")
		return nil
	},
}

var cleanCmd = &cmds.Command{
	Arguments: []cmds.Argument{
		{
			Name: "path",
		},
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		repo, _ := fsrepo.Open(request.Arguments[0])
		pstore, _ := pstoreds.NewPeerstore(context.Background(), repo.Datastore(), pstoreds.DefaultOpts())
		address := pstore.PeersWithAddrs()
		for _, add := range address {
			pstore.ClearAddrs(add)
		}
		return nil
	},
}
