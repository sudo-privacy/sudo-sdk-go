package sudoclient

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"context"
	"sync/atomic"

	"fmt"
)

// 参考 googlee.golang.org/api/transport/grpc/pool.go
// 移除singlePool
// roundRobinConnPool增加tokenSource成员，用于释放

// ConnPool is a pool of grpc.ClientConns.
type ConnPool interface {
	// Conn returns a ClientConn from the pool.
	//
	// Conns aren't returned to the pool.
	Conn() *grpc.ClientConn

	// Num returns the number of connections in the pool.
	//
	// It will always return the same value.
	Num() int

	// Close closes every ClientConn in the pool.
	//
	// The error returned by Close may be a single error or multiple errors.
	Close() error

	// ConnPool implements grpc.ClientConnInterface to enable it to be used directly with generated proto stubs.
	grpc.ClientConnInterface
}

type roundRobinConnPool struct {
	conns []*grpc.ClientConn

	idx uint32 // access via sync/atomic

	tokenSource TokenSource
}

func (p *roundRobinConnPool) Num() int {
	return len(p.conns)
}

func (p *roundRobinConnPool) Conn() *grpc.ClientConn {
	i := atomic.AddUint32(&p.idx, 1)
	return p.conns[i%uint32(len(p.conns))]
}

func (p *roundRobinConnPool) Close() error {
	err := p.tokenSource.Close()
	if err != nil {
		return err
	}
	var errs multiError
	for _, conn := range p.conns {
		if err := conn.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func (p *roundRobinConnPool) Invoke(
	ctx context.Context,
	method string,
	args, reply interface{},
	opts ...grpc.CallOption,
) error {
	return p.Conn().Invoke(ctx, method, args, reply, opts...)
}

func (p *roundRobinConnPool) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	return p.Conn().NewStream(ctx, desc, method, opts...)
}

// multiError represents errors from multiple conns in the group.
//
// TODO: figure out how and whether this is useful to export. End users should
// not be depending on the transport/grpc package directly, so there might need
// to be some service-specific multi-error type.
type multiError []error

func (m multiError) Error() string {
	s, n := "", 0
	for _, e := range m {
		if e != nil {
			if n == 0 {
				s = e.Error()
			}
			n++
		}
	}
	switch n {
	case 0:
		return "(0 errors)"
	case 1:
		return s
	case 2:
		return s + " (and 1 other error)"
	}
	return fmt.Sprintf("%s (and %d other errors)", s, n-1)
}

func dialPool(ctx context.Context, opts ...DialOption) (ConnPool, error) {
	o := newDefaultdialOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	if o.grpcConnPool != nil {
		return o.grpcConnPool, nil
	}

	if o.tokenSource == nil {
		tokenSource, err := NewUserAccountToken(ctx, opts...)
		if err != nil {
			return nil, err
		}
		o.tokenSource = tokenSource
	}

	poolSize := o.grpcConnPoolSize
	if poolSize <= 1 {
		poolSize = 1
	}

	pool := &roundRobinConnPool{
		tokenSource: o.tokenSource,
	}
	for i := 0; i < poolSize; i++ {
		conn, err := dial(ctx, o)
		if err != nil {
			pool.Close() // NOTE: error from Close is ignored.
			return nil, err
		}
		pool.conns = append(pool.conns, conn)
	}
	return pool, nil
}

func dial(ctx context.Context, o *dialOptions) (*grpc.ClientConn, error) {
	if o.grpcEndPoint == "" {
		return nil, errors.New("endpoint not set")
	}
	if o.tokenSource == nil {
		return nil, errors.New("tokenSource not set")
	}
	perGrpcCred := grpc.WithPerRPCCredentials(
		o.tokenSource,
	)
	grpcOpts := o.grpcOpts
	grpcOpts = append(grpcOpts, perGrpcCred)
	return grpc.DialContext(ctx, o.grpcEndPoint, grpcOpts...)
}
