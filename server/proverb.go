package server

import (
	"context"

	"github.com/go-openapi/strfmt"
	goproverbs "github.com/indiependente/go-proverbs"
	service "github.com/indiependente/go-proverbs-server/rpc/service/v1"
	grpctransport "github.com/indiependente/go-proverbs-server/transport/grpc"
)

// Proverb returns a Go proverb.
func (srv Server) Proverb(ctx context.Context, req *service.ProverbRequest) (*service.ProverbResponse, error) {
	id, ok := ctx.Value(grpctransport.RequestIDKey).(string)
	if !ok {
		srv.logger.Error("could not type assert request id from context", nil)
	}

	prvb := goproverbs.Proverb()

	srv.logger.Proverb(prvb).RequestID(strfmt.UUID(id)).Info("proverb generated")
	return &service.ProverbResponse{
		Proverb: prvb,
	}, nil
}
