package core

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseQueryMediaOrderIsStable(t *testing.T) {
	query := url.Values{
		"audio": {""},
		"video": {""},
	}

	medias := ParseQuery(query)
	require.Len(t, medias, 2)
	require.Equal(t, KindVideo, medias[0].Kind)
	require.Equal(t, KindAudio, medias[1].Kind)
}
