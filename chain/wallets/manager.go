package wallets

import (
	"sync"

	"github.com/eosforce/relay/cmd/config"
	"github.com/fanyang1988/eos-go"
	"github.com/pkg/errors"
)

// WalletInfo info to wallet
type WalletInfo struct {
	APIURL     string
	URL        string
	WalletName string
	ChainName  string
}

// NewWalletInfoFromCfg create a wallet info from cfg
func NewWalletInfoFromCfg(cfg *config.ChainCfg) WalletInfo {
	return WalletInfo{
		APIURL:     cfg.ApiURL,
		URL:        cfg.URL,
		WalletName: "default",
		ChainName:  cfg.Name,
	}
}

type manager struct {
	infos map[string]WalletInfo // chainName to walletInfo
	mtx   sync.RWMutex
}

// GetSigner get a signer for push trx
func (m *manager) GetSigner(chainName string) (eos.Signer, error) {
	info, err := m.GetInfo(chainName)
	if err != nil {
		return nil, err
	}
	return eos.NewWalletSigner(eos.New(info.URL), info.WalletName), nil
}

// GetInfo get wallet info
func (m *manager) GetInfo(chainName string) (WalletInfo, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	info, ok := m.infos[chainName]
	if !ok {
		return WalletInfo{}, errors.New("no chainName found")
	}

	return info, nil
}

func (m *manager) RegWallet(w WalletInfo) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.infos[w.ChainName] = w
}

var mng *manager

func Get() *manager {
	return mng
}

func init() {
	mng = &manager{
		infos: make(map[string]WalletInfo, 32),
	}
}
