package backend

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogDBPath(t *testing.T) {
	base := t.TempDir()
	chainIDStr := "42984928492928428424243444"
	chainID, ok := new(big.Int).SetString(chainIDStr, 10)
	require.True(t, ok)
	expected := filepath.Join(base, "subdir", chainIDStr, "log.db")
	path, err := prepLogDBPath(chainID, filepath.Join(base, "subdir"))
	require.NoError(t, err)
	require.Equal(t, expected, path)

	// Check it still works when directories exist
	require.NoError(t, os.WriteFile(path, []byte("test"), 0o644))

	path, err = prepLogDBPath(chainID, filepath.Join(base, "subdir"))
	require.NoError(t, err)
	require.Equal(t, expected, path)
}
