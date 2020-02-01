package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// ContextKey represents a unique key used to store/retrieve values from a context.Context
type ContextKey int

const (
	// RequestIDKey is the key to use to retrieve the ID of a request from a context.Context
	RequestIDKey ContextKey = iota
)

// ServerInterceptor tracing unary interceptor function to handle request tagging per RPC call.
func ServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("could not generate new UUID: %w", err)
	}

	ctx = context.WithValue(ctx, RequestIDKey, id.String())

	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}
