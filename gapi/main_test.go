package gapi

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	db "github.com/stuartfranke/golang-backend-master-class/db/sqlc"
	"github.com/stuartfranke/golang-backend-master-class/token"
	"github.com/stuartfranke/golang-backend-master-class/util"
	"github.com/stuartfranke/golang-backend-master-class/worker"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, role string, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(username, role, duration)
	require.NoError(t, err)

	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}

	return metadata.NewIncomingContext(context.Background(), md)
}