package paintApi

import (
	"context"
	"path"
	"time"

	"paint/internal/def"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func unaryServerLogger(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (_ interface{}, err error) {
	log := newLogger(ctx, info.FullMethod)
	ctx = structlog.NewContext(ctx, log)
	return handler(ctx, req)
}

func streamServerLogger(
	srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	ctx := stream.Context()
	log := newLogger(ctx, info.FullMethod)
	ctx = structlog.NewContext(ctx, log)
	wrapped := grpc_middleware.WrapServerStream(stream)
	wrapped.WrappedContext = ctx
	return handler(srv, wrapped)
}

func newLogger(ctx context.Context, fullMethod string) *structlog.Logger {
	kvs := []interface{}{
		def.LogFunc, path.Base(fullMethod),
		def.LogGRPCCode, "",
	}
	if p, ok := peer.FromContext(ctx); ok {
		kvs = append(kvs, def.LogRemote, p.Addr.String())
	}
	return structlog.New(kvs...)
}

func unaryServerRecover(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (_ interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			log := structlog.FromContext(ctx, nil)
			log.PrintErr("panic", def.LogGRPCCode, codes.Internal, "err", p,
				structlog.KeyStack, structlog.Auto)
			err = status.Errorf(codes.Internal, "%v", p)
		}
	}()
	return handler(ctx, req)
}

func streamServerRecover(
	srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if p := recover(); p != nil {
			log := structlog.FromContext(stream.Context(), nil)
			log.PrintErr("panic", "err", p, structlog.KeyStack, structlog.Auto)
			err = status.Errorf(codes.Internal, "%v", p)
		}
	}()
	return handler(srv, stream)
}

func unaryServerAccessLog(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (_ interface{}, err error) {
	resp, err := handler(ctx, req)
	log := structlog.FromContext(ctx, nil)
	log.SetDefaultKeyvals(structlog.KeyTime, time.Now().Format(time.StampMicro))
	logHandler(log, err)
	return resp, err
}

func streamServerAccessLog(
	srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	log := structlog.FromContext(stream.Context(), nil)
	log.SetDefaultKeyvals(structlog.KeyTime, time.Now().Format(time.StampMicro))
	log.Info("started")
	err = handler(srv, stream)
	logHandler(log, err)
	return err
}

func logHandler(log *structlog.Logger, err error) {
	s := status.Convert(err)
	code, msg := s.Code(), s.Message()
	switch code {
	case codes.Unknown:
		log.PrintErr("failed to handle", def.LogGRPCCode, code, "err", msg)
	case codes.InvalidArgument:
		log.Warn("handled", def.LogGRPCCode, code, "err", msg)
	case codes.DeadlineExceeded:
		log.Warn("handled", def.LogGRPCCode, code)
	case codes.NotFound:
		log.Info("handled", def.LogGRPCCode, code, "err", msg)
	case codes.AlreadyExists:
		log.Info("handled", def.LogGRPCCode, code, "err", msg)
	case codes.PermissionDenied:
		log.Warn("handled", def.LogGRPCCode, code, "err", msg)
	case codes.ResourceExhausted:
		log.Info("handled", def.LogGRPCCode, code, "err", msg)
	case codes.FailedPrecondition:
		log.Info("handled", def.LogGRPCCode, code, "err", msg)
	case codes.Aborted:
		log.Info("handled", def.LogGRPCCode, code, "err", msg)
	case codes.OutOfRange:
		log.Warn("handled", def.LogGRPCCode, code, "err", msg)
	case codes.Unimplemented:
		log.PrintErr("failed to handle", def.LogGRPCCode, code, "err", msg)
	case codes.Internal:
		log.PrintErr("failed to handle", def.LogGRPCCode, code, "err", msg)
	case codes.Unavailable:
		log.Warn("handled", def.LogGRPCCode, code, "err", msg)
	case codes.DataLoss:
		log.PrintErr("failed to handle", def.LogGRPCCode, code, "err", msg)
	case codes.Unauthenticated:
		log.Warn("handled", def.LogGRPCCode, code)
	case codes.OK, codes.Canceled:
	}
}
