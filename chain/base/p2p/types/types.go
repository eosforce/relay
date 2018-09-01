package p2pTypes

// P2PMessage
type P2PMessage interface {
	GetP2PMsg() interface{}
	GetType() byte
}

type MessageHandler func(msg P2PMessage)
