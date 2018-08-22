package chainMsg

import (
	"testing"
	"time"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/eos-handler"
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

	watcher := startWatcher(func(action eosHandler.ActionData) {
		seelog.Infof("action %v %v %v", action.Action.Name, action.Action.Account, action.Action.Authorization)
		if action.Action.Name != "transfer" || action.Action.Account != "eosio" {
			return
		}
		msg, err := builder.Build(&action)
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

func startWatcher(f func(action eosHandler.ActionData)) *eosHandler.EosWatcher {
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

	watcher := eosHandler.NewEosWatcher("eosforce1", apiURL, p2pAddrs)
	watcher.RegHandler(f)
	err := watcher.Start()
	if err != nil {
		return nil
	}

	return watcher
}
