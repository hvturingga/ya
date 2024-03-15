//go:build windows
// +build windows

package embed

import _ "embed"

//go:embed ya-daemon-windows-amd64.exe
var Daemon []byte
