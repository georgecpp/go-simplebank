package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/georgecpp/go-simplebank/db/sqlc"
	"github.com/georgecpp/go-simplebank/token"
	"github.com/georgecpp/go-simplebank/util"
	"github.com/georgecpp/go-simplebank/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
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

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, duration time.Duration) context.Context {
	ctx := context.Background()
	accessToken, _, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string {
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(ctx, md)
}