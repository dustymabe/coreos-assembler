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

package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/mdlayher/vsock"
)

const (
	// defaultVsockPort is the default SSH port for vsock connections.
	defaultVsockPort = 22
)

// HybridDialer is a Dialer that handles both vsock and TCP addresses.
// When the address starts with "vsock:", it dials via AF_VSOCK.
// Otherwise, it delegates to an embedded RetryDialer for standard TCP.
type HybridDialer struct {
	*RetryDialer

	// VsockRetries is the number of retries for vsock connections.
	// Defaults to DefaultRetries if zero.
	VsockRetries int

	// VsockTimeout is the timeout for each vsock dial attempt.
	// Defaults to DefaultTimeout if zero.
	VsockTimeout time.Duration
}

// NewHybridDialer creates a HybridDialer with default settings for
// both vsock and TCP connections.
func NewHybridDialer() *HybridDialer {
	return &HybridDialer{
		RetryDialer:  NewRetryDialer(),
		VsockRetries: DefaultRetries,
		VsockTimeout: DefaultTimeout,
	}
}

// parseVsockAddress parses a vsock address string of the form "vsock:<CID>"
// or "vsock:<CID>:<port>". Returns the CID and port.
func parseVsockAddress(address string) (uint32, uint32, error) {
	// Strip the "vsock:" prefix
	addr := strings.TrimPrefix(address, "vsock:")

	parts := strings.SplitN(addr, ":", 2)
	cid, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid vsock CID %q: %v", parts[0], err)
	}

	port := uint32(defaultVsockPort)
	if len(parts) == 2 {
		p, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid vsock port %q: %v", parts[1], err)
		}
		port = uint32(p)
	}

	return uint32(cid), port, nil
}

// Dial connects to the given address. If the address starts with "vsock:",
// it connects via AF_VSOCK. Otherwise, it delegates to the TCP RetryDialer.
func (d *HybridDialer) Dial(network, address string) (net.Conn, error) {
	if strings.HasPrefix(address, "vsock:") {
		return d.dialVsock(address)
	}
	return d.RetryDialer.Dial(network, address)
}

// dialVsock connects to a vsock address with retry logic.
func (d *HybridDialer) dialVsock(address string) (c net.Conn, err error) {
	cid, port, err := parseVsockAddress(address)
	if err != nil {
		return nil, err
	}

	retries := d.VsockRetries
	if retries == 0 {
		retries = DefaultRetries
	}

	for i := 0; i < retries; i++ {
		c, err = vsock.Dial(cid, port, nil)
		if err == nil {
			return
		}
	}
	return
}
