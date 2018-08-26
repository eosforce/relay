package orderBook

import "sort"

type SortedOrderList struct {
	orderList  []Order
	sorted     bool
	isBigFirst bool
}

// NewSortedOrderList make a list init space by cap
func NewSortedOrderList(isBigFirst bool, cap int) SortedOrderList {
	return SortedOrderList{
		orderList:  make([]Order, 0, cap),
		sorted:     true,
		isBigFirst: isBigFirst,
	}
}

const orderCapLen = 64 * 1024

// NewBuyOrderList new buy order list, max price is in last then first process
func NewBuyOrderList() SortedOrderList {
	return NewSortedOrderList(true, orderCapLen)
}

// NewSellOrderList new sell order list, min price is in last then first process
func NewSellOrderList() SortedOrderList {
	return NewSortedOrderList(false, orderCapLen)
}

// Len is the number of elements in the collection.
func (s *SortedOrderList) Len() int {
	return len(s.orderList)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s *SortedOrderList) Less(i, j int) bool {
	// note the max is on the tail, this good for pop success exchange
	if s.orderList[i].Price == s.orderList[j].Price {
		return s.orderList[i].UpdateTime.UnixNano() > s.orderList[j].UpdateTime.UnixNano()
	}
	if s.isBigFirst {
		return s.orderList[i].Price.Less(&s.orderList[j].Price)
	} else {
		return s.orderList[i].Price.More(&s.orderList[j].Price)
	}
}

// Swap swaps the elements with indexes i and j.
func (s *SortedOrderList) Swap(i, j int) {
	if i == j {
		return
	}
	t := s.orderList[i]
	s.orderList[i] = s.orderList[j]
	s.orderList[j] = t
}

// Sort sort if is no sorted
func (s *SortedOrderList) Sort() {
	if s.sorted {
		return
	}

	sort.Sort(s)
	s.sorted = true
}

// Append append new order
func (s *SortedOrderList) Append(order Order) {
	s.orderList = append(s.orderList, order)
	s.sorted = false
}

// Pop pop the last (which is max to use) order
func (s *SortedOrderList) Pop() Order {
	if len(s.orderList) == 0 {
		return Order{}
	}

	// sort first if no sorted
	s.Sort()

	l := len(s.orderList) - 1
	res := s.orderList[l]
	s.orderList = s.orderList[:l]
	return res
}

// Last the last (which is max to use) order
func (s *SortedOrderList) Last() *Order {
	if len(s.orderList) == 0 {
		return nil
	}

	// sort first if no sorted
	s.Sort()

	l := len(s.orderList) - 1
	return &s.orderList[l]
}
