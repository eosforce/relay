package eosforce

import eos "github.com/fanyang1988/eos-go"

var actionFee map[eos.ActionName]eos.Asset

func init() {
	actionFee = map[eos.ActionName]eos.Asset{
		//eosio.bios
		eos.ActN("newaccount"): eos.NewEOSAsset(1000),
		eos.ActN("updateauth"): eos.NewEOSAsset(1000),
		eos.ActN("deleteauth"): eos.NewEOSAsset(1000),

		//System
		eos.ActN("transfer"):     eos.NewEOSAsset(100),
		eos.ActN("vote"):         eos.NewEOSAsset(500),
		eos.ActN("unfreeze"):     eos.NewEOSAsset(100),
		eos.ActN("claim"):        eos.NewEOSAsset(300),
		eos.ActN("updatebp"):     eos.NewEOSAsset(100 * 10000),
		eos.ActN("setemergency"): eos.NewEOSAsset(10 * 10000),

		//eosio.token
		eos.ActN("issue"):  eos.NewEOSAsset(100),
		eos.ActN("create"): eos.NewEOSAsset(10 * 10000),

		//eosio.msig
		eos.ActN("propose"):   eos.NewEOSAsset(1000),
		eos.ActN("approve"):   eos.NewEOSAsset(1000),
		eos.ActN("unapprove"): eos.NewEOSAsset(1000),
		eos.ActN("cancel"):    eos.NewEOSAsset(1000),
		eos.ActN("exec"):      eos.NewEOSAsset(1000),
	}
}

// GetFeeByAction get fee by action name
func GetFeeByAction(actionName eos.ActionName) eos.Asset {
	fee, ok := actionFee[actionName]
	//fmt.Printf("get key %v %v %v %v", fee.String(), actionName, actionFee, ok)
	if ok {
		//fmt.Printf("get key %v", fee.String())
		return fee
	} else {
		return eos.Asset{Amount: 0, Symbol: eos.EOSSymbol}
	}
}

// GetFeeByActions get fee sum by actions
func GetFeeByActions(actions []*eos.Action) eos.Asset {
	res := eos.Asset{Amount: 0, Symbol: eos.EOSSymbol}
	for _, act := range actions {
		feeAct := GetFeeByAction(act.Name)
		res = res.Add(feeAct)
		//fmt.Printf("add key %v %v\n", res.String(), feeAct.String())
	}
	return res
}
