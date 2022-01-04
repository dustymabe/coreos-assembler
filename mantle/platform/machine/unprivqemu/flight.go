// Copyright 2019 Red Hat
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

package unprivqemu

import (
	"github.com/coreos/pkg/capnslog"

	"github.com/coreos/mantle/platform"
	"github.com/coreos/mantle/platform/conf"
)

const (
	Platform platform.Name = "qemu"
)

// Options contains QEMU-specific options for the flight.
type Options struct {
	// DiskImage is the full path to the disk image to boot in QEMU.
	DiskImage string
	// DiskSize if non-empty will expand the disk
	DiskSize string
	Board    string
	Firmware string
	Memory   string

	NbdDisk       bool
	MultiPathDisk bool
	Native4k      bool
	Nvme          bool

	//Option to create a temporary software TPM - true by default
	Swtpm bool

	*platform.Options
}

type flight struct {
	*platform.BaseFlight
	opts *Options
}

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/mantle", "platform/machine/qemu")
)

func NewFlight(opts *Options) (platform.Flight, error) {
	bf, err := platform.NewBaseFlight(opts.Options, Platform)
	if err != nil {
		return nil, err
	}

	qf := &flight{
		BaseFlight: bf,
		opts:       opts,
	}

	return qf, nil
}

func (af *flight) ConfigTooLarge(ud conf.UserData) bool {

	// Config is not too large
	return false
}

// NewCluster creates a Cluster instance, suitable for running virtual
// machines in QEMU.
func (qf *flight) NewCluster(rconf *platform.RuntimeConfig) (platform.Cluster, error) {
	bc, err := platform.NewBaseCluster(qf.BaseFlight, rconf)
	if err != nil {
		return nil, err
	}

	qc := &Cluster{
		BaseCluster: bc,
		flight:      qf,
	}

	qf.AddCluster(qc)

	return qc, nil
}
