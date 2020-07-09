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

package hostutil

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/utils/mount"
)

// FakeHostUtil is a fake HostUtils implementation for testing
type FakeHostUtil struct {
	MountPoints []mount.MountPoint
	Filesystem  map[string]hostutil.FileType

	mutex sync.Mutex
}

// NewFakeHostUtil returns a struct that implements the HostUtils interface
// for testing
// TODO: no callers were initializing the struct with any MountPoints. Check
// if those are still being used by any callers and if MountPoints still need
// to be a part of the struct.
func NewFakeHostUtil(fs map[string]hostutil.FileType) *FakeHostUtil {
	return &FakeHostUtil{
		Filesystem: fs,
	}
}

// Compile-time check to make sure FakeHostUtil implements interface
var _ hostutil.HostUtils = &FakeHostUtil{}

// getDeviceNameFromMountLinux find the device name from /proc/mounts in which
// the mount path reference should match the given plugin mount directory. In case no mount path reference
// matches, returns the volume name taken from its given mountPath
func getDeviceNameFromMount(mounter mount.Interface, mountPath, pluginMountDir string) (string, error) {
	refs, err := mounter.GetMountRefs(mountPath)
	if err != nil {
		klog.V(4).Infof("GetMountRefs failed for mount path %q: %v", mountPath, err)
		return "", err
	}
	if len(refs) == 0 {
		klog.V(4).Infof("Directory %s is not mounted", mountPath)
		return "", fmt.Errorf("directory %s is not mounted", mountPath)
	}
	for _, ref := range refs {
		if strings.HasPrefix(ref, pluginMountDir) {
			volumeID, err := filepath.Rel(pluginMountDir, ref)
			if err != nil {
				klog.Errorf("Failed to get volume id from mount %s - %v", mountPath, err)
				return "", err
			}
			return volumeID, nil
		}
	}

	return path.Base(mountPath), nil
}

// DeviceOpened checks if block device referenced by pathname is in use by
// checking if is listed as a device in the in-memory mountpoint table.
func (hu *FakeHostUtil) DeviceOpened(pathname string) (bool, error) {
	hu.mutex.Lock()
	defer hu.mutex.Unlock()

	for _, mp := range hu.MountPoints {
		if mp.Device == pathname {
			return true, nil
		}
	}
	return false, nil
}

// PathIsDevice always returns true
func (hu *FakeHostUtil) PathIsDevice(pathname string) (bool, error) {
	return true, nil
}

// GetDeviceNameFromMount given a mount point, find the volume id
func (hu *FakeHostUtil) GetDeviceNameFromMount(mounter mount.Interface, mountPath, pluginMountDir string) (string, error) {
	return getDeviceNameFromMount(mounter, mountPath, pluginMountDir)
}

// MakeRShared checks if path is shared and bind-mounts it as rshared if needed.
// No-op for testing
func (hu *FakeHostUtil) MakeRShared(path string) error {
	return nil
}

// GetFileType checks for file/directory/socket/block/character devices.
// Defaults to Directory if otherwise unspecified.
func (hu *FakeHostUtil) GetFileType(pathname string) (hostutil.FileType, error) {
	if t, ok := hu.Filesystem[pathname]; ok {
		return t, nil
	}
	return hostutil.FileType("Directory"), nil
}

// PathExists checks if pathname exists.
func (hu *FakeHostUtil) PathExists(pathname string) (bool, error) {
	if _, ok := hu.Filesystem[pathname]; ok {
		return true, nil
	}
	return false, nil
}

// EvalHostSymlinks returns the path name after evaluating symlinks.
// No-op for testing
func (hu *FakeHostUtil) EvalHostSymlinks(pathname string) (string, error) {
	return pathname, nil
}

// GetOwner returns the integer ID for the user and group of the given path
// Not implemented for testing
func (hu *FakeHostUtil) GetOwner(pathname string) (int64, int64, error) {
	return -1, -1, errors.New("GetOwner not implemented")
}

// GetSELinuxSupport tests if pathname is on a mount that supports SELinux.
// Not implemented for testing
func (hu *FakeHostUtil) GetSELinuxSupport(pathname string) (bool, error) {
	return false, errors.New("GetSELinuxSupport not implemented")
}

// GetMode returns permissions of pathname.
// For volcano simulator testing
func (hu *FakeHostUtil) GetMode(pathname string) (os.FileMode, error) {
	return os.FileMode(0755), nil
}
