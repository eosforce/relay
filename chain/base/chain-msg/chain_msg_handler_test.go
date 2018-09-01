package chainMsg

import (
	"testing"
	"time"

	"github.com/eosforce/relay/const"

	"github.com/eosforce/relay/chain/base/handler"
	"github.com/eosforce/relay/types"

	"github.com/cihub/seelog"
)

func TestHandler(t *testing.T) {
	defer seelog.Flush()

	hh := NewChainMsgHandler()
	hh.AddHandler("main", "biosbpa", func(msg *ChainMsg) {
		seelog.Infof("on msg a1 %v", *msg)
	})
	hh.AddHandler("main", "biosbpa", func(msg *ChainMsg) {
		seelog.Infof("on msg a2 %v", *msg)
	})
	hh.AddHandler("main", "biosbpb", func(msg *ChainMsg) {
		seelog.Infof("on msg b %v", *msg)
	})

	hh.Start()

	builder := Transfer2MsgBuilder{}

	watcher := startWatcher(func(action types.ActionData) {
		seelog.Infof("action %v %v", action.Name, action.Account)
		if action.Name != "transfer" || action.Account != "eosio" {
			return
		}
		msg, err := builder.BuildFromEOSForce(&action)
		if err != nil {
			seelog.Errorf("build msg err By %s", err.Error())
			return
		}

		if msg == nil {
			seelog.Error("build msg err nil")
			return
		}

		seelog.Tracef("push msg %v %v", msg.MsgTyp, msg)
		hh.PushMsg(msg)
	})

	time.Sleep(20 * time.Second)

	watcher.Stop()
	watcher.Wait()

	hh.Stop()
	hh.Wait()
}

func startWatcher(f func(action types.ActionData)) *handler.EosWatcher {
	// start a chain http api :8889, p2p address 9001
	apiURL := "http://127.0.0.1:8890"
	p2pAddrs := []string{
		"127.0.0.1:9002",
		"127.0.0.1:9003",
		"127.0.0.1:9004",
		"127.0.0.1:9005",
		"127.0.0.1:9006",
		"127.0.0.1:9007",
		"127.0.0.1:9008",
		"127.0.0.1:9009",
		"127.0.0.1:9010",
	}

	watcher := handler.NewEosWatcher(consts.TypeBaseEosforce, "eosforce1", apiURL, p2pAddrs)
	watcher.RegHandler(f)
	err := watcher.Start()
	if err != nil {
		return nil
	}

	return watcher
}
