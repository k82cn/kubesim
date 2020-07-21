package config

import (
	"bytes"
	"io/ioutil"

	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// SinkConfig sink configration
type SinkConfig struct {
	Sink      string            `yaml:"sink"`
	Parameter map[string]string `yaml:"parameter"`
}

// SinkConfigFromYaml parse sink configuration from yaml
func SinkConfigFromYaml(path string) (*SinkConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return SinkConfigFromBytes(data)
}

// SinkConfigFromBytes parse sink configuration from byte
func SinkConfigFromBytes(data []byte) (*SinkConfig, error) {
	c := &SinkConfig{}

	reader := bytes.NewReader(data)
	decoder := k8syaml.NewYAMLToJSONDecoder(reader)
	err := decoder.Decode(c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
