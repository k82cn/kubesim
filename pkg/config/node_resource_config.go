package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

// NodeResourceConfig contains all node resource configuration
type NodeResourceConfig struct {
	NodeClasses []NodeClasses `yaml:"nodeclasses"`
}

// NodeClasses contains
type NodeClasses struct {
	Name      string            `yaml:"name"`
	Labels    map[string]string `yaml:"labels"`
	Resources nodeResource      `yaml:"resources"`
}

type nodeResource struct {
	Capacity map[string]string `yaml:"capacity"`
}

// NodeConfigFromYaml parse node configuration from yaml
func NodeConfigFromYaml(path, resourceName string) (*NodeClasses, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NodeConfigFromBytes(data, resourceName)
}

// NodeConfigFromBytes parse node configuration from byte
func NodeConfigFromBytes(data []byte, resourceName string) (*NodeClasses, error) {
	c := &NodeResourceConfig{}

	reader := bytes.NewReader(data)
	decoder := k8syaml.NewYAMLToJSONDecoder(reader)
	err := decoder.Decode(c)

	if err != nil {
		return nil, err
	}

	// Validate pod class names for uniquenes
	classNames := map[string]bool{}
	for _, class := range c.NodeClasses {
		name := strings.ToLower(class.Name)
		if _, exists := classNames[name]; exists {
			return nil, fmt.Errorf("node class name [%s] is not unique", name)
		}
		classNames[name] = true
	}

	for _, class := range c.NodeClasses {
		name := strings.ToLower(class.Name)
		if name == resourceName {
			return &class, nil
		}
	}
	return nil, err
}
