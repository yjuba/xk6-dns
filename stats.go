package xk6_dns

import (
	"context"
	"fmt"
	"time"

	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/metrics"
	"go.k6.io/k6/stats"
)

var (
	DialCount = stats.New("dns.dial.count", stats.Counter)
	DialError = stats.New("dns.dial.error", stats.Counter)

	RequestCount = stats.New("dns.request.count", stats.Counter)
	RequestError = stats.New("dns.request.error", stats.Counter)

	ResponseTime = stats.New("dns.response.time", stats.Trend, stats.Time)
)

func reportDial(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: DialCount,
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func reportDialError(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: DialError,
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func reportRequest(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: RequestCount,
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func reportRequestError(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: RequestError,
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func reportResponseTime(ctx context.Context, rtt time.Duration) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: ResponseTime,
		Time:   now,
		Value:  float64(rtt.Milliseconds()),
	})

	return nil
}

func reportDataSent(ctx context.Context, value float64) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: metrics.DataSent,
		Time:   now,
		Value:  value,
	})

	return nil
}

func reportDataReceived(ctx context.Context, value float64) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: metrics.DataReceived,
		Time:   now,
		Value:  value,
	})

	return nil
}
