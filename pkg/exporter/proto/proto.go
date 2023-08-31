package proto

import (
	"context"
)

//go:generate protoc --go_out=. ./inspector.proto

type ProbeType string

var (
	ProbeTypeMetrics ProbeType = "metrics"
	ProbeTypeEvent   ProbeType = "event"
)

type RawEvent struct {
	Netns     uint32
	EventType string
	EventBody string
}

type Probe interface {
	Start(ctx context.Context, probeType ProbeType)
	Close(probeType ProbeType) error
	Ready() bool
	Name() string
}

type MetricProbe interface {
	Probe
	GetMetricNames() []string
	Collect(ctx context.Context) (map[string]map[uint32]uint64, error)
}

type EventProbe interface {
	Probe
	GetEventNames() []string
	Register(receiver chan<- RawEvent) error
}
