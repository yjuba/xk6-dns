package xk6_dns

import (
	"context"
	"fmt"
	"time"

	"github.com/miekg/dns"
)

type K6DNS struct {
	client *dns.Client

	Version string
}

func NewK6DNS(version string) *K6DNS {
	return &K6DNS{
		client:  &dns.Client{},
		Version: version,
	}
}

func (k *K6DNS) SetDialTimeout(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	k.client.DialTimeout = d
	return nil
}

func (k *K6DNS) SetReadTimeout(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	k.client.ReadTimeout = d
	return nil
}

func (k *K6DNS) SetWriteTimeout(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	k.client.WriteTimeout = d
	return nil
}

func (k *K6DNS) Resolve(ctx context.Context, addr, query, qtypeStr string) (string, error) {
	qtype, ok := dns.StringToType[qtypeStr]
	if !ok {
		return "", fmt.Errorf("unknown query type: %s", qtypeStr)
	}

	msg := &dns.Msg{}
	msg.Id = dns.Id()
	msg.RecursionDesired = true
	msg.Question = make([]dns.Question, 1)
	msg.Question[0] = dns.Question{
		Name:   query,
		Qtype:  qtype,
		Qclass: dns.ClassINET,
	}

	reportDial(ctx)
	conn, err := NewK6UDPConn(addr)
	if err != nil {
		reportDialError(ctx)
		return err.Error(), nil
	}
	defer func() {
		conn.Close()
		reportDataReceived(ctx, float64(conn.rxBytes))
		reportDataSent(ctx, float64(conn.txBytes))
	}()

	reportRequest(ctx)
	resp, rtt, err := k.client.ExchangeWithConn(msg, &dns.Conn{Conn: conn})
	if err != nil {
		reportRequestError(ctx)
		return err.Error(), nil
	}
	reportResponseTime(ctx, rtt)

	return resp.String(), nil
}
