/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"strconv"

	"github.com/google/cadvisor/events"
	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kubernetes/pkg/kubelet/cadvisor"

	simulatorconfig "volcano.sh/simulator/pkg/config"
)

// New new fake cadvisor.Interface
func New(nodeName string, nodeClass *simulatorconfig.NodeClasses) *Fake {
	numCores := fakeNumCores
	memoryCapacity := fakeMemoryCapacity

	if nodeClass != nil {
		if v, ok := nodeClass.Resources.Capacity[string(v1.ResourceCPU)]; ok {
			numCores, _ = strconv.Atoi(v)
		}
		if v, ok := nodeClass.Resources.Capacity[string(v1.ResourceMemory)]; ok {
			a := resource.MustParse(v)
			memoryCapacity64, _ := a.AsInt64()
			memoryCapacity = int(memoryCapacity64)
		}
	}

	return &Fake{
		NodeName:       nodeName,
		NumCores:       numCores,
		MemoryCapacity: memoryCapacity,
	}
}

// Fake cadvisor.Interface implementation.
type Fake struct {
	NodeName       string
	NumCores       int
	MemoryCapacity int
}

const (
	// FakeKernelVersion is a fake kernel version for testing.
	FakeKernelVersion = "3.16.0-0.bpo.4-amd64"
	// FakeContainerOSVersion is a fake OS version for testing.
	FakeContainerOSVersion = "Debian GNU/Linux 7 (wheezy)"

	fakeNumCores       = 1
	fakeMemoryCapacity = 4026531840
	fakeDockerVersion  = "1.13.1"
)

var _ cadvisor.Interface = new(Fake)

// Start is a fake implementation of Interface.Start.
func (c *Fake) Start() error {
	return nil
}

// ContainerInfo is a fake implementation of Interface.ContainerInfo.
func (c *Fake) ContainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (*cadvisorapi.ContainerInfo, error) {
	return new(cadvisorapi.ContainerInfo), nil
}

// ContainerInfoV2 is a fake implementation of Interface.ContainerInfoV2.
func (c *Fake) ContainerInfoV2(name string, options cadvisorapiv2.RequestOptions) (map[string]cadvisorapiv2.ContainerInfo, error) {
	return map[string]cadvisorapiv2.ContainerInfo{}, nil
}

// SubcontainerInfo is a fake implementation of Interface.SubcontainerInfo.
func (c *Fake) SubcontainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (map[string]*cadvisorapi.ContainerInfo, error) {
	return map[string]*cadvisorapi.ContainerInfo{}, nil
}

// DockerContainer is a fake implementation of Interface.DockerContainer.
func (c *Fake) DockerContainer(name string, req *cadvisorapi.ContainerInfoRequest) (cadvisorapi.ContainerInfo, error) {
	return cadvisorapi.ContainerInfo{}, nil
}

// MachineInfo is a fake implementation of Interface.MachineInfo.
func (c *Fake) MachineInfo() (*cadvisorapi.MachineInfo, error) {
	// Simulate a machine with 1 core and 3.75GB of memory.
	// We set it to non-zero values to make non-zero-capacity machines in Kubemark.
	return &cadvisorapi.MachineInfo{
		NumCores:       c.NumCores,
		InstanceID:     cadvisorapi.InstanceID(c.NodeName),
		MemoryCapacity: uint64(c.MemoryCapacity),
	}, nil
}

// VersionInfo is a fake implementation of Interface.VersionInfo.
func (c *Fake) VersionInfo() (*cadvisorapi.VersionInfo, error) {
	return &cadvisorapi.VersionInfo{
		KernelVersion:      FakeKernelVersion,
		ContainerOsVersion: FakeContainerOSVersion,
		DockerVersion:      fakeDockerVersion,
	}, nil
}

// ImagesFsInfo is a fake implementation of Interface.ImagesFsInfo.
func (c *Fake) ImagesFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

// RootFsInfo is a fake implementation of Interface.RootFsInfo.
func (c *Fake) RootFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

// WatchEvents is a fake implementation of Interface.WatchEvents.
func (c *Fake) WatchEvents(request *events.Request) (*events.EventChannel, error) {
	return new(events.EventChannel), nil
}

// GetDirFsInfo is a fake implementation of Interface.GetDirFsInfo.
func (c *Fake) GetDirFsInfo(path string) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}
