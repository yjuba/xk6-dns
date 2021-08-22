package xk6_dns

import (
	"context"
	"fmt"
	"time"

	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/metrics"
	"go.k6.io/k6/stats"
)

func countError(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: stats.New("dns.error.count", stats.Counter),
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func countDial(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: stats.New("dns.dial.count", stats.Counter),
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func countRequest(ctx context.Context) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: stats.New("dns.request.count", stats.Counter),
		Time:   now,
		Value:  float64(1),
	})

	return nil
}

func countResponseRTT(ctx context.Context, rtt time.Duration) error {
	state := lib.GetState(ctx)
	if state == nil {
		return fmt.Errorf("state is nil")
	}

	now := time.Now()
	stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
		Metric: stats.New("dns.response.rtt", stats.Trend, stats.Time),
		Time:   now,
		Value:  float64(rtt.Milliseconds()),
	})

	return nil
}

func countDataSent(ctx context.Context, value float64) error {
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

func countDataReceived(ctx context.Context, value float64) error {
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
