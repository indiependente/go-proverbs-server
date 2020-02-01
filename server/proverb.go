package server

import (
	"context"

	"github.com/go-openapi/strfmt"
	goproverbs "github.com/indiependente/go-proverbs"
	"github.com/indiependente/go-proverbs-server/rpc"
	grpctransport "github.com/indiependente/go-proverbs-server/transport/grpc"
)

// Proverb returns a Go proverb.
func (srv Server) Proverb(ctx context.Context, req *rpc.ProverbRequest) (*rpc.ProverbResponse, error) {
	id, ok := ctx.Value(grpctransport.RequestIDKey).(string)
	if !ok {
		srv.logger.Error("could not type assert request id from context", nil)
	}

	prvb := goproverbs.Proverb()

	srv.logger.Proverb(prvb).RequestID(strfmt.UUID(id)).Info("proverb generated")
	return &rpc.ProverbResponse{
		Proverb: prvb,
	}, nil
}
