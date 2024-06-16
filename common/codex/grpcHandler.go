package codex

import (
	"context"
	"fmt"
	"go-zero-mall/common/codex/types"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Error(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func Errorf(code Code, format string, args ...interface{}) *Status {
	code.msg = fmt.Sprintf(format, args...)
	return Error(code)
}

func FromProto(pbMsg proto.Message) CodeX {
	msg, ok := pbMsg.(*types.Status)
	if ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Code{code: int(msg.Code)}
		}
		return &Status{sts: msg}
	}

	return Errorf(ServerErr, "invalid proto message get %v", pbMsg)
}

func (s *Status) Proto() *types.Status {
	return s.sts
}

func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}

func FromCode(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func GetGrpcStatusFromError(err error) *status.Status {
	// 获取错误的根本原因
	err = errors.Cause(err)

	// 检查错误是否是 CodeX 类型
	if code, ok := err.(CodeX); ok {
		// 调用 gRPCStatusFromCodeX 函数将 CodeX 转换为 gRPC status
		grpcStatus, e := gRPCStatusFromCodeX(code)
		if e == nil {
			// 如果没有错误，返回转换后的 gRPC status
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	// 根据具体的错误类型返回相应的 gRPC status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromCodeX(Canceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromCodeX(Deadline)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}

func gRPCStatusFromCodeX(code CodeX) (*status.Status, error) {
	var sts *Status

	// 使用 switch 语句根据传入的 code 类型进行不同处理
	switch v := code.(type) {
	case *Status:
		// 如果 code 是 *Status 类型，直接将其赋值给 sts
		sts = v
	case Code:
		// 如果 code 是 Code 类型，使用 FromCode 函数将其转换为 *Status 类型
		sts = FromCode(v)
	default:
		// 如果 code 是其他类型，将其转换为 *Status
		sts = FromCode(Code{code.Code(), code.Message()})
		// 遍历 code 的 Details，并将其添加到 sts 的 Details 中
		for _, detail := range code.Details() {
			if msg, ok := detail.(proto.Message); ok {
				_, _ = sts.WithDetails(msg)
			}
		}
	}

	stas := status.New(codes.Unknown, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.Proto())
}

func String(s string) Code {
	if len(s) == 0 {
		return OK
	}
	code, err := strconv.Atoi(s)
	if err != nil {
		return ServerErr
	}

	return Code{code: code}
}

func toXCode(grpcStatus *status.Status) Code {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return Deadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

func GrpcStatusToXCode(gstatus *status.Status) CodeX {
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			return FromProto(pb)
		}
	}

	return toXCode(gstatus)
}
