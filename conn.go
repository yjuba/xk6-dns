package xk6_dns

import (
	"net"
	"sync/atomic"
	"time"
)

type k6UDPConn struct {
	*net.UDPConn

	rxBytes int64
	txBytes int64
}

func NewK6UDPConn(addr string) (*k6UDPConn, error) {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	c, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return nil, err
	}
	return &k6UDPConn{UDPConn: c}, nil
}

func (k *k6UDPConn) GetRXBytes() int64 {
	return k.rxBytes
}

func (k *k6UDPConn) GetTXBytes() int64 {
	return k.txBytes
}

func (k *k6UDPConn) Read(b []byte) (int, error) {
	n, err := k.UDPConn.Read(b)
	if n > 0 {
		atomic.AddInt64(&k.rxBytes, int64(n))
	}
	return n, err
}

func (k *k6UDPConn) Write(b []byte) (int, error) {
	n, err := k.UDPConn.Write(b)
	if n > 0 {
		atomic.AddInt64(&k.txBytes, int64(n))
	}
	return n, err
}

func (k *k6UDPConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return k.UDPConn.ReadFrom(p)
}

func (k *k6UDPConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return k.UDPConn.WriteTo(p, addr)
}

func (k *k6UDPConn) Close() error {
	return k.UDPConn.Close()
}

func (k *k6UDPConn) LocalAddr() net.Addr {
	return k.UDPConn.LocalAddr()
}

func (k *k6UDPConn) SetDeadline(t time.Time) error {
	return k.UDPConn.SetDeadline(t)
}

func (k *k6UDPConn) SetReadDeadline(t time.Time) error {
	return k.UDPConn.SetDeadline(t)
}

func (k *k6UDPConn) SetWriteDeadline(t time.Time) error {
	return k.UDPConn.SetWriteDeadline(t)
}

var _ net.PacketConn = &k6UDPConn{}
