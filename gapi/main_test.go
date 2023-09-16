package gapi

import (
	"testing"
	"time"

	db "github.com/georgecpp/go-simplebank/db/sqlc"
	"github.com/georgecpp/go-simplebank/util"
	"github.com/georgecpp/go-simplebank/worker"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	
	return server
}