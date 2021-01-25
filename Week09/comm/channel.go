package comm

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net"
	"sync"
)

// Channel used by message pusher send msg to write goroutine.
type Channel struct {
	conn   net.Conn
	Writer *bufio.Writer
	Reader *bufio.Reader
	signal chan *Proto
	mutex  sync.RWMutex
}

func (ch *Channel) Read() {
	p := &Proto{}
	p.Read(ch.Reader)
	log.Printf("收到消息%v", p)
	p.Body = []byte("{\"code\":\"0000\"}")
	log.Printf("发送消息%v", p)
	ch.signal <- p
}

func (ch *Channel) Write() {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()
	proto := <-ch.signal
	proto.Write(ch.Writer)
}

func NewChannnel(conn net.Conn) *Channel {
	c := new(Channel)
	c.Writer = bufio.NewWriter(conn)
	c.Reader = bufio.NewReader(conn)
	c.signal = make(chan *Proto, 10)
	return c
}
