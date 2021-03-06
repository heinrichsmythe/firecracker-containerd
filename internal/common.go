// Copyright 2018-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package internal

const (
	// StdinPort represents vsock port to be used for stdin
	StdinPort = 11000
	// StdoutPort represents vsock port to be used for stdout
	StdoutPort = 11001
	// StderrPort represents vsock port to be used for stderr
	StderrPort = 11002
	// DefaultBufferSize represents buffer size in bytes to used for IO between runtime and agent
	DefaultBufferSize = 1024

	// FirecrackerSockName is the name of the Firecracker VMM API socket
	FirecrackerSockName = "firecracker.sock"
	// FirecrackerLogFifoName is the name of the Firecracker VMM log FIFO
	FirecrackerLogFifoName = "fc-logs.fifo"
	// FirecrackerMetricsFifoName is the name of the Firecracker VMM metrics FIFO
	FirecrackerMetricsFifoName = "fc-metrics.fifo"

	// TODO these strings are hardcoded throughout the containerd codebase, it may
	// be worth sending them a PR to define them as constants accessible to shim
	// implementations like our own

	// ShimAddrFileName is the name of the file in which a shim's API address can be found, used by
	// containerd to reconnect after it restarts
	ShimAddrFileName = "address"

	// ShimLogFifoName is the name of the FIFO created by containerd for a shim to write its logs to
	ShimLogFifoName = "log"

	// OCIConfigName is the name of the OCI bundle's config field
	OCIConfigName = "config.json"

	// BundleRootfsName is the name of the bundle's directory for holding the container's rootfs
	BundleRootfsName = "rootfs"

	// VMIDEnvVarKey is the environment variable key used to provide a VMID to a shim process
	VMIDEnvVarKey = "FIRECRACKER_VM_ID"

	// FCSocketFDEnvKey is the environment variable key used to provide the FD of the fccontrol listening socket to a shim
	FCSocketFDEnvKey = "FCCONTROL_SOCKET_FD"

	// ShimBinaryName is the name of the runtime shim binary
	ShimBinaryName = "containerd-shim-aws-firecracker"
)
