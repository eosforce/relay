package orderBook

import (
	"sync"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/types"
)

// OnExchangeCallback when buy and sell order make a exchange, call this
type OnExchangeCallback func(book *OrderBook, buy, sell *Order) error

// orderCmd command to send to OrderBook
type orderCmd struct {
	order   Order
	resChan chan error
}

// OrderBook a order book contain two order list which for buy and sell and last exchange price
type OrderBook struct {
	currPrice int64
	buyList   SortedOrderList
	sellList  SortedOrderList

	wg        sync.WaitGroup
	orderChan chan orderCmd

	pair types.ExchangePair
	cb   OnExchangeCallback
}

// NewOrderBook create a order book for pair
func NewOrderBook(pair types.ExchangePair, cb OnExchangeCallback) *OrderBook {
	return &OrderBook{
		currPrice: 0,
		buyList:   NewBuyOrderList(),
		sellList:  NewSellOrderList(),
		pair:      pair,
		cb:        cb,
		orderChan: make(chan orderCmd, 1024),
	}
}

// GetCurrPrice get curr price
func (o *OrderBook) GetCurrPrice() int64 {
	// TODO add mutex
	return o.currPrice
}

func (o *OrderBook) pushOrder(order *Order) error {
	resChan := make(chan error, 1)
	o.orderChan <- orderCmd{
		order:   *order,
		resChan: resChan,
	}
	err := <-resChan
	return err
}

// Buy add a buy order
func (o *OrderBook) Buy(account types.Account, asset types.Asset, price int64) error {
	// FIXME buy other
	order := NewBuy(account, o.pair, asset, price)
	return o.pushOrder(&order)
}

// Sell add a buy order
func (o *OrderBook) Sell(account types.Account, asset types.Asset, price int64) error {
	// FIXME buy other
	order := NewSell(account, o.pair, asset, price)
	return o.pushOrder(&order)
}

// Start
func (o *OrderBook) Start() error {
	o.wg.Add(1)
	go func() {
		defer o.wg.Done()
		for {
			isStop, err := o.Loop()
			if err != nil {
				seelog.Errorf("err By %s", err.Error())
			}
			if isStop {
				seelog.Warnf("orderbook %s stop", o)
				return
			}
		}
	}()
	return nil
}

// Stop
func (o *OrderBook) Stop() {
	close(o.orderChan)
}

// Wait
func (o *OrderBook) Wait() {
	o.wg.Wait()
}

// Loop if err return err, if is stop return true, nil
func (o *OrderBook) Loop() (bool, error) {
	cmd, ok := <-o.orderChan
	if !ok {
		return true, nil
	}

	// process cmd
	seelog.Tracef("process order %s", cmd.order)
	err := o.processOrder(&cmd.order)

	if cmd.resChan != nil {
		cmd.resChan <- err
	}

	return false, err
}
