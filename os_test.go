package hello

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsUserDir(t *testing.T) {
	var userHomeDir, _ = os.UserHomeDir()
	PrintKeyValue("user home dir", userHomeDir)
	var userCacheDir, _ = os.UserCacheDir()
	PrintKeyValue("user cache dir", userCacheDir)
	var userConfigDir, _ = os.UserConfigDir()
	PrintKeyValue("user config dir", userConfigDir)
}

func TestOsEnv(t *testing.T) {
	PrintKeyValue("shell", os.Getenv("SHELL"))
}

func TestOsHostname(t *testing.T) {
	var hostname, _ = os.Hostname()
	PrintKeyValue("hostname", hostname)
}

func TestOsExec(t *testing.T) {
	var cmd = exec.Command("go", "version")
	PrintKeyValue("cmd", cmd)
	if out, err := cmd.CombinedOutput(); err != nil {
		assert.Failf(t, "Failed to run command", "err=%v", err)
	} else {
		PrintKeyValue("cmd out+err", string(out))
	}
	PrintKeyValue("cmd pid", cmd.ProcessState.Pid())
	PrintKeyValue("cmd success", cmd.ProcessState.Success())
	PrintKeyValue("cmd exitcode", cmd.ProcessState.ExitCode())
	PrintKeyValue("cmd usertime", cmd.ProcessState.UserTime())
	PrintKeyValue("cmd systime", cmd.ProcessState.SystemTime())
}
