package sudoclient

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type dialOptions struct {
	accountSetting
	cc               *grpc.ClientConn
	grpcEndPoint     string
	grpcConnPool     ConnPool
	grpcConnPoolSize int
	grpcOpts         []grpc.DialOption
	httpEndPoint     string
	tokenSource      TokenSource
	tokenTimeout     time.Duration
}

type DialOption interface {
	Apply(*dialOptions)
}

func newDefaultdialOptions() *dialOptions {
	return &dialOptions{
		grpcEndPoint: "",
		httpEndPoint: "",
		grpcOpts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		tokenTimeout: time.Minute * 5,
	}
}

func WithTokenSource(tokenSource TokenSource) DialOption {
	return withTokenSource{tokenSource}
}

type withTokenSource struct {
	tokenSource TokenSource
}

func (w withTokenSource) Apply(o *dialOptions) {
	o.tokenSource = w.tokenSource
}

func WithConnPool(grpcConnPool ConnPool) DialOption {
	return withConnPool{grpcConnPool}
}

type withConnPool struct{ grpcConnPool ConnPool }

func (w withConnPool) Apply(o *dialOptions) {
	o.grpcConnPool = w.grpcConnPool
}

// WithGrpcEndpoint set grpc service endpoint
func WithGrpcEndpoint(endPoint string) DialOption {
	return withGrpcEndpoint{endPoint}
}

type withGrpcEndpoint struct{ endPoint string }

func (w withGrpcEndpoint) Apply(o *dialOptions) {
	o.grpcEndPoint = w.endPoint
}

func WithHTTPEndpoint(endpoint string) DialOption {
	return withHTTPEndpoint{endpoint}
}

type withHTTPEndpoint struct {
	endpoint string
}

func (w withHTTPEndpoint) Apply(o *dialOptions) {
	o.httpEndPoint = w.endpoint
}

func WithGrpcClientInterceptor(interceptor grpc.UnaryClientInterceptor) DialOption {
	return withGrpcClientInterceptor{interceptor}
}

type withGrpcClientInterceptor struct {
	interceptor grpc.UnaryClientInterceptor
}

func (w withGrpcClientInterceptor) Apply(o *dialOptions) {
	o.grpcOpts = append(o.grpcOpts, grpc.WithChainUnaryInterceptor(w.interceptor))
}

func WithAccount(account, password string) DialOption {
	return withAccountSetting{
		account: accountSetting{
			Account:  account,
			Password: password,
		},
	}
}

type withAccountSetting struct {
	account accountSetting
}

func (w withAccountSetting) Apply(o *dialOptions) {
	o.accountSetting = w.account
}
