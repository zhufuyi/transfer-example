// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	transferV1 "transfer/api/transfer/v1"
	"transfer/internal/ecode"
	"transfer/internal/rpcclient"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		transferV1.RegisterTransferServer(server, NewTransferServer())
	})
}

var _ transferV1.TransferServer = (*transfer)(nil)

type transfer struct {
	transferV1.UnimplementedTransferServer
}

// NewTransferServer create a server
func NewTransferServer() transferV1.TransferServer {
	rpcclient.InitDtmServerResolver()
	return &transfer{}
}

// Transfer 转账
func (s *transfer) Transfer(ctx context.Context, req *transferV1.TransferRequest) (*transferV1.TransferReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	var (
		// 直连ip方式 或 服务发现方式，由配置文件决定
		dtmServer      = rpcclient.GetDtmEndpoint()
		transferServer = rpcclient.GetTransferEndpoint()
	)

	logger.Debug("server endpoint", logger.String("dtm", dtmServer), logger.String("transfer", transferServer))

	transOutData := &transferV1.TransOutRequest{
		Amount: req.Amount,
		UserId: req.FromUserId,
	}
	transInData := &transferV1.TransInRequest{
		Amount: req.Amount,
		UserId: req.ToUserId,
	}
	gid := dtmgrpc.MustGenGid(dtmServer)
	m := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(transferServer+"/api.transfer.v1.transfer/TransOut", transOutData).
		Add(transferServer+"/api.transfer.v1.transfer/TransIn", transInData)
	m.WaitResult = true
	err = m.Submit()
	if err != nil {
		logger.Error("Transfer error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}
	return &transferV1.TransferReply{}, nil
}

// TransOut 转出
func (s *transfer) TransOut(ctx context.Context, req *transferV1.TransOutRequest) (*transferV1.TransOutReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	logger.Info("转出成功", logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	return &transferV1.TransOutReply{}, nil
}

// TransIn 转入
func (s *transfer) TransIn(ctx context.Context, req *transferV1.TransInRequest) (*transferV1.TransInReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	logger.Info("转入成功", logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	return &transferV1.TransInReply{}, nil
}