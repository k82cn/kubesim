package metrics

import (
	"github.com/volcano-sh/kubesim/pkg/metrics/config"
	"k8s.io/klog"
)

// GlogSink implements metrics.Interface
type GlogSink struct {
	sink *config.SinkConfig
	node string
}

// NewGlogSink create a new GlogSink
func NewGlogSink(sink *config.SinkConfig, node string) Interface {
	return &GlogSink{
		sink: sink,
		node: node,
	}
}

// Initialization is used to initialize GlogSink
func (ms *GlogSink) Initialization() error {
	klog.Infoln("GlogSink.Initialization() done")
	return nil
}

// LogNodeMetrics is used to log node metrics by glog
func (ms *GlogSink) LogNodeMetrics(nm *NodeMetric) error {
	switch nm.MetricType {
	case "static":
		klog.Infof("node metric, type %s, name %s, cpu %s, memory %s", nm.MetricType, ms.node, nm.Capacity["cpu"], nm.Capacity["memory"])
	case "real":
		klog.Infof("node metric, type %s, name %s, cpu %s, memory %s, timestamp %d", nm.MetricType, ms.node, nm.Capacity["cpu"], nm.Capacity["memory"], nm.SampleTime.Unix())
	default:
		klog.Errorf("unsupported node metrics type %s", nm.MetricType)
	}
	return nil
}

// LogVolcanlJobMetrics is used to log volcano job metrics by glog
func (ms *GlogSink) LogVolcanlJobMetrics() error {
	return nil
}

// LogPodMetrics is used to log pod metrics by glog
func (ms *GlogSink) LogPodMetrics() error {
	return nil
}
