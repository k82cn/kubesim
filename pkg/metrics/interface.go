package metrics

import (
	"time"

	"volcano.sh/kubesim/pkg/metrics/config"
)

// Interface is the interface used to output cluster metrics
type Interface interface {
	// Initialization is used to initialize the interface
	Initialization() error

	// LogNodeMetrics is used to log node metrics
	LogNodeMetrics(nm *NodeMetric) error

	// LogNodeMetrics is used to log node metrics
	LogVolcanlJobMetrics() error

	// LogNodeMetrics is used to log node metrics
	LogPodMetrics() error
}

// NodeMetric define the node metric
type NodeMetric struct {
	MetricType string // static/real
	Capacity   map[string]string
	SampleTime time.Time
}

// VolcanlJobMetric define the volcano job metric
type VolcanlJobMetric struct {
}

// PodMetric define the pod metric
type PodMetric struct {
}

// ManufactureSink will manufacture a sink according to
func ManufactureSink(config *config.SinkConfig, node string) Interface {
	sinkType := ""
	if config != nil {
		sinkType = config.Sink
	}
	switch sinkType {
	case "mysql":
		return NewMysqlSink(config, node)
	case "log":
		return NewGlogSink(config, node)
	default:
		return NewGlogSink(config, node)
	}
}
