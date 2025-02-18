package token

import (
	"context"

	tokenproto "github.com/linera-hacker/linera-dapps/service/kline/proto/kline/zeus/v1/token"
	token "github.com/linera-hacker/linera-dapps/service/kline/zeus/pkg/mw/v1/token"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteToken(ctx context.Context, in *tokenproto.DeleteTokenRequest) (*tokenproto.DeleteTokenResponse, error) {
	handler, err := token.NewHandler(
		ctx,
		token.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteToken",
			"In", in,
			"Error", err,
		)
		return &tokenproto.DeleteTokenResponse{}, status.Error(codes.Internal, "internal server err")
	}

	info, err := handler.DeleteToken(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteToken",
			"In", in,
			"Error", err,
		)
		return &tokenproto.DeleteTokenResponse{}, status.Error(codes.Internal, "internal server err")
	}

	return &tokenproto.DeleteTokenResponse{Info: info}, nil
}
