package gocmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModule(t *testing.T) {
	t.Parallel()

	gocmd, err := NewGoCmd()
	require.NoError(t, err)

	module, err := gocmd.Module(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "github.com/mt-sre/go-ci", module)
}

// TestTidyConfig_Option tests the Option method of the
// TidyConfig struct, which allows setting configuration options
func TestTidyConfig_Option(t *testing.T) {
	config := &TidyConfig{}

	withGoVersion := func(version string) TidyOption {
		return tidyOptionFunc(func(c *TidyConfig) {
			c.GoVersion = version
		})
	}

	withWorkingDir := func(dir string) TidyOption {
		return tidyOptionFunc(func(c *TidyConfig) {
			c.WorkingDir = dir
		})
	}

	config.Option(withGoVersion("1.16.4"), withWorkingDir("/tmp"))

	if config.GoVersion != "1.16.4" {
		t.Errorf("expected Go version to be %q, got %q", "1.16.4", config.GoVersion)
	}

	if config.WorkingDir != "/tmp" {
		t.Errorf("expected WorkingDir to be %q, got %q", "/tmp", config.WorkingDir)
	}
}

type tidyOptionFunc func(*TidyConfig)

func (f tidyOptionFunc) ConfigureTidy(config *TidyConfig) {
	f(config)
}

func TestNewGoCmd(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		cmd, err := NewGoCmd()

		require.NoError(err, "expected no error")
		assert.NotEmpty(cmd.cfg.BinPath, "expected default BinPath to be set")

	})

	t.Run("WithOptions", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		expectedBinPath := "/path/to/go"

		cmd, err := NewGoCmd(WithBinPath(expectedBinPath))

		require.NoError(err, "expected no error")
		assert.Equal(expectedBinPath, cmd.cfg.BinPath, "expected BinPath to be %q", expectedBinPath)
	})
}
