package latestver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLatestVer(t *testing.T) {
	module := "github.com/hashicorp/raft"
	latestVer, err := LatestVersion(module)

	require.Nil(t, err)
	require.Equal(t, "1.1.2", latestVer)
}
