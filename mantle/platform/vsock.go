// Copyright 2025 Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package platform

import (
	"fmt"
	"math/rand"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	// ioctlVhostVsockSetGuestCID is the ioctl number for VHOST_VSOCK_SET_GUEST_CID.
	// This ioctl is used to atomically check and reserve a vsock CID on the host.
	ioctlVhostVsockSetGuestCID = 0x4008AF60

	// vhostVsockDevicePath is the path to the vhost-vsock device.
	vhostVsockDevicePath = "/dev/vhost-vsock"

	// minVsockCID is the minimum usable CID (0=hypervisor, 1=local, 2=host).
	minVsockCID = 3

	// maxVsockCID is the maximum usable CID.
	maxVsockCID = 0xFFFFFFFE

	// vsockCIDAttempts is the number of CID candidates to try before giving up.
	vsockCIDAttempts = 128
)

// VsockAvailable checks whether the host supports vsock by verifying
// that /dev/vhost-vsock exists and is accessible.
func VsockAvailable() bool {
	fd, err := unix.Open(vhostVsockDevicePath, unix.O_RDWR|unix.O_CLOEXEC|unix.O_NONBLOCK, 0)
	if err != nil {
		plog.Warningf("vsock not available (%s): %v", vhostVsockDevicePath, err)
		return false
	}
	unix.Close(fd)
	return true
}

// vsockCIDInUse probes whether a CID is already in use by attempting the
// VHOST_VSOCK_SET_GUEST_CID ioctl on the given vhost-vsock fd.
// Returns true if the CID is in use (EADDRINUSE), false if it is free.
func vsockCIDInUse(fd int, cid uint32) (bool, error) {
	// The ioctl expects a pointer to a uint64 containing the CID.
	val := uint64(cid)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), ioctlVhostVsockSetGuestCID, uintptr(unsafe.Pointer(&val)))
	if errno != 0 {
		if errno == unix.EADDRINUSE {
			return true, nil
		}
		return false, fmt.Errorf("ioctl VHOST_VSOCK_SET_GUEST_CID failed: %v", errno)
	}
	return false, nil
}

// FindUnusedVsockCID opens /dev/vhost-vsock and finds an unused CID by
// probing random candidates with the VHOST_VSOCK_SET_GUEST_CID ioctl.
// On success, it returns the CID. Note that there is a small TOCTOU race
// between this check and QEMU reserving the CID, but with randomly chosen
// CIDs from a ~4 billion range, collisions are vanishingly unlikely.
func FindUnusedVsockCID() (uint32, error) {
	fd, err := unix.Open(vhostVsockDevicePath, unix.O_RDWR|unix.O_CLOEXEC|unix.O_NONBLOCK, 0)
	if err != nil {
		return 0, fmt.Errorf("opening %s: %v", vhostVsockDevicePath, err)
	}
	defer unix.Close(fd)

	for i := 0; i < vsockCIDAttempts; i++ {
		cid := uint32(rand.Int63n(int64(maxVsockCID-minVsockCID+1))) + minVsockCID

		inUse, err := vsockCIDInUse(fd, cid)
		if err != nil {
			return 0, err
		}
		if !inUse {
			plog.Debugf("Allocated vsock CID %d", cid)
			return cid, nil
		}
	}

	return 0, fmt.Errorf("failed to find unused vsock CID after %d attempts", vsockCIDAttempts)
}
