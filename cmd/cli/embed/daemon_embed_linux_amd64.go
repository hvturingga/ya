//go:build linux
// +build linux

package embed

import _ "embed"

//go:embed ya-daemon-linux-amd64
var Daemon []byte
