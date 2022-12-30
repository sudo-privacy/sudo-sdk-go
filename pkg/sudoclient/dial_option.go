package sudoclient

import (
	"net/http"
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
	httpTransport    http.RoundTripper
	tokenSource      TokenSource
	tokenTimeout     time.Duration
}

// DialOption 设置连接相关配置，比如 endpoint、账号、密码。
type DialOption interface {
	Apply(*dialOptions)
}

func newDefaultDialOptions() *dialOptions {
	return &dialOptions{
		grpcEndPoint: "",
		httpEndPoint: "",
		grpcOpts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		httpTransport: http.DefaultTransport,
		tokenTimeout:  time.Minute * 5,
	}
}

// WithTokenSource 设置tokenSource。
func WithTokenSource(tokenSource TokenSource) DialOption {
	return withTokenSource{tokenSource}
}

type withTokenSource struct {
	tokenSource TokenSource
}

func (w withTokenSource) Apply(o *dialOptions) {
	o.tokenSource = w.tokenSource
}

// WithConnPool 设置连接池。
func WithConnPool(grpcConnPool ConnPool) DialOption {
	return withConnPool{grpcConnPool}
}

type withConnPool struct{ grpcConnPool ConnPool }

func (w withConnPool) Apply(o *dialOptions) {
	o.grpcConnPool = w.grpcConnPool
}

// WithGrpcEndpoint 设置gRPC服务endpoint。
func WithGrpcEndpoint(endPoint string) DialOption {
	return withGrpcEndpoint{endPoint}
}

type withGrpcEndpoint struct{ endPoint string }

func (w withGrpcEndpoint) Apply(o *dialOptions) {
	o.grpcEndPoint = w.endPoint
}

// WithGrpcEndpoint 设置HTTP endpoint。
func WithHTTPEndpoint(endpoint string) DialOption {
	return withHTTPEndpoint{endpoint}
}

type withHTTPEndpoint struct {
	endpoint string
}

func (w withHTTPEndpoint) Apply(o *dialOptions) {
	o.httpEndPoint = w.endpoint
}

// WithGrpcClientInterceptor 增加gRPC拦截器。
func WithGrpcClientInterceptor(interceptor grpc.UnaryClientInterceptor) DialOption {
	return withGrpcClientInterceptor{interceptor}
}

type withGrpcClientInterceptor struct {
	interceptor grpc.UnaryClientInterceptor
}

func (w withGrpcClientInterceptor) Apply(o *dialOptions) {
	o.grpcOpts = append(o.grpcOpts, grpc.WithChainUnaryInterceptor(w.interceptor))
}

// WithAccount 设置账号、密码。
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

// WithGrpcDialOption 增加grpc.DialOption 设置
func WithGrpcDialOption(grpcOpt grpc.DialOption) DialOption {
	return withGrpcDialOption{
		dialOption: grpcOpt,
	}
}

type withGrpcDialOption struct {
	dialOption grpc.DialOption
}

func (w withGrpcDialOption) Apply(o *dialOptions) {
	o.grpcOpts = append(o.grpcOpts, w.dialOption)
}

// WithHTTPTransport 修改http Client Transport 设置
func WithHTTPTransport(transport http.RoundTripper) DialOption {
	return withHTTPTransport{transport: transport}
}

type withHTTPTransport struct {
	transport http.RoundTripper
}

func (w withHTTPTransport) Apply(o *dialOptions) {
	o.httpTransport = w.transport
}
