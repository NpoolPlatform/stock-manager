// +build !codeanalysis

package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/stock-manager/pkg/const"
	crud "github.com/NpoolPlatform/stock-manager/pkg/crud/stock"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npoolcommon "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/stockmgr"

	"github.com/google/uuid"

	"google.golang.org/protobuf/types/known/structpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateStock(ctx context.Context, in *npool.CreateStockRequest) (*npool.CreateStockResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		logger.Sugar().Errorf("invalid request good id: %v", err)
		return &npool.CreateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create stock: %v", err)
		return &npool.CreateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStockResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateStocks(ctx context.Context, in *npool.CreateStocksRequest) (*npool.CreateStocksResponse, error) {
	for _, info := range in.GetInfos() {
		if _, err := uuid.Parse(info.GetGoodID()); err != nil {
			logger.Sugar().Errorf("invalid request good id: %v", err)
			return &npool.CreateStocksResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateStocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := schema.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create stocks: %v", err)
		return &npool.CreateStocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStocksResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateStock(ctx context.Context, in *npool.UpdateStockRequest) (*npool.UpdateStockResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		logger.Sugar().Errorf("invalid request good id: %v", err)
		return &npool.UpdateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorf("invalid stock id: %v", err)
		return &npool.UpdateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update stock: %v", err)
		return &npool.UpdateStockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateStockResponse{
		Info: info,
	}, nil
}

func stockFieldsToFields(fields map[string]*structpb.Value) (map[string]interface{}, error) {
	newFields := map[string]interface{}{}

	for k, v := range fields {
		switch k {
		case constant.FieldID:
			newFields[k] = v.GetStringValue()
		case constant.StockFieldGoodID:
			newFields[k] = v.GetStringValue()
		case constant.StockFieldTotal:
			newFields[k] = uint32(v.GetNumberValue())
		case constant.StockFieldInService:
			newFields[k] = uint32(v.GetNumberValue())
		case constant.StockFieldSold:
			newFields[k] = uint32(v.GetNumberValue())
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}
	return newFields, nil
}

func (s *Server) UpdateStockFields(ctx context.Context, in *npool.UpdateStockFieldsRequest) (*npool.UpdateStockFieldsResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("invalid stock id: %v", err)
		return &npool.UpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	fields, err := stockFieldsToFields(in.GetFields())
	if err != nil {
		logger.Sugar().Errorf("invalid stock fields: %v", err)
		return &npool.UpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(fields) == 0 {
		logger.Sugar().Errorf("empty stock fields: %v", err)
		return &npool.UpdateStockFieldsResponse{}, status.Error(codes.Internal, "empty stock fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.UpdateFields(ctx, id, fields)
	if err != nil {
		logger.Sugar().Errorf("fail update stock: %v", err)
		return &npool.UpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateStockFieldsResponse{
		Info: info,
	}, nil
}

func (s *Server) AtomicUpdateStockFields(ctx context.Context, in *npool.AtomicUpdateStockFieldsRequest) (*npool.AtomicUpdateStockFieldsResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("invalid stock id: %v", err)
		return &npool.AtomicUpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	fields, err := stockFieldsToFields(in.GetFields())
	if err != nil {
		logger.Sugar().Errorf("invalid stock fields: %v", err)
		return &npool.AtomicUpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(fields) == 0 {
		logger.Sugar().Errorf("empty stock fields: %v", err)
		return &npool.AtomicUpdateStockFieldsResponse{}, status.Error(codes.Internal, "empty stock fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.AtomicUpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.AtomicUpdateFields(ctx, id, fields)
	if err != nil {
		logger.Sugar().Errorf("fail atomic update stock: %v", err)
		return &npool.AtomicUpdateStockFieldsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AtomicUpdateStockFieldsResponse{
		Info: info,
	}, nil
}

func stockCondsToConds(conds map[string]*npoolcommon.FilterCond) (map[string]*cruder.Cond, error) {
	newConds := map[string]*cruder.Cond{}

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}

		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.StockFieldGoodID:
			newConds[k] = &cruder.Cond{
				Op:  v.Op,
				Val: v.Val.GetStringValue(),
			}
		case constant.StockFieldTotal:
			fallthrough //nolint
		case constant.StockFieldInService:
			fallthrough //nolint
		case constant.StockFieldSold:
			newConds[k] = &cruder.Cond{
				Op:  v.Op,
				Val: v.Val.GetNumberValue(),
			}
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}

	return newConds, nil
}

func (s *Server) ExistStockConds(ctx context.Context, in *npool.ExistStockCondsRequest) (*npool.ExistStockCondsResponse, error) {
	conds, err := stockCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid stock fields: %v", err)
		return &npool.ExistStockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(conds) == 0 {
		logger.Sugar().Errorf("empty stock fields: %v", err)
		return &npool.ExistStockCondsResponse{}, status.Error(codes.Internal, "empty stock fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistStockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.ExistConds(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail check stock: %v", err)
		return &npool.ExistStockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStockCondsResponse{
		Result: exist,
	}, nil
}
