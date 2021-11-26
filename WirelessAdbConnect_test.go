package main

import (
	"github.com/Ericwyn/GoTools/shell"
	"testing"
)

func TestAdbConnect(t *testing.T) {
	shell.RunShellRes("adb", "connect", "192.168.199.222:39845")
}
