package protocol

import (
	"net"
	"sync"
)

type NetHandler interface {
	// HandleRequest should get all data from r, process it and write result to w
	HandleRequest(
		r *net.TCPConn,
		w *net.TCPConn,
		connCh chan *net.TCPConn,
		addrCh chan string,
		netRequest NetRequest,
		isInboundConn bool,
		originalDst string) *net.TCPConn
	// HandleResponse should get all data from r, process it and write result to w
	HandleResponse(r *net.TCPConn, w *net.TCPConn, netRequest NetRequest, isInboundConn bool, forceClose bool)
}

var bufferPool = sync.Pool{
	New: func() interface{} { return make([]byte, 0xffff) },
}
