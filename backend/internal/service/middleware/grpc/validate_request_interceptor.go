package grpc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
)

type Validatable interface {
	Validate() error
}

func (s Service) ValidateRequestInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if r, ok := req.Any().(Validatable); ok {
				if err := r.Validate(); err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation error: %v", err))
				}
			} else {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("request type does not implement Validatable interface"))
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
