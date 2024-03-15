//go:build linux
// +build linux

package embed

import _ "embed"

//go:embed ya-daemon-linux-arm64
var Daemon []byte
