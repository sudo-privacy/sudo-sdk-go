package sudoclient

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/jwt"
	basiccommon "gitlab.sudoprivacy.cn/weixy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/service/common"
)

// TokenSource 是对 [credentials.PerRPCCredentials]的简单封装，每次访问可能触发token刷新。
type TokenSource interface {
	credentials.PerRPCCredentials
	Close() error
}

type accountSetting struct {
	Account  string
	Password string
}

type userAccountToken struct {
	basiccommon.CommonClient
	accessToken string
	opts        *dialOptions

	expiry       time.Time
	refreshToken string
	mutex        sync.RWMutex
}

// NewUserAccountToken 根据用户名、密码创建tokenSource，endpoint、账号等需要通过opts配置。
func NewUserAccountToken(ctx context.Context, opts ...DialOption) (TokenSource, error) {
	o := newDefaultDialOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	if o.cc == nil {
		cc, err := grpc.DialContext(ctx, o.grpcEndPoint, o.grpcOpts...)
		if err != nil {
			return nil, err
		}
		o.cc = cc
	}

	token := &userAccountToken{
		CommonClient: basiccommon.NewCommonClient(o.cc),
		opts:         o,
	}

	err := token.init(ctx)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (source *userAccountToken) tokenWithContext(ctx context.Context) (*oauth2.Token, error) {
	token, err := source.token()
	if err != nil || !token.Valid() {
		refreshErr := source.refresh(ctx)
		if refreshErr != nil {
			return nil, refreshErr
		}
		return source.token()
	}
	return token, nil
}

func (source *userAccountToken) token() (*oauth2.Token, error) {
	source.mutex.RLock()
	defer source.mutex.RUnlock()
	return &oauth2.Token{
		AccessToken:  source.accessToken,
		RefreshToken: source.refreshToken,
		Expiry:       source.expiry,
	}, nil
}

func (source *userAccountToken) Close() error {
	return source.opts.cc.Close()
}

func (source *userAccountToken) init(ctx context.Context) error {
	resp, err := source.CommonClient.CreateJwt(ctx, &jwt.CreateJwtRequest{
		Account:  source.opts.Account,
		Password: source.opts.Password,
	})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("create jwt failed,account:%s", source.opts.Account))
	}
	source.mutex.Lock()
	defer source.mutex.Unlock()
	source.accessToken = resp.Access
	source.refreshToken = resp.Refresh
	source.expiry = time.Now().Add(source.opts.tokenTimeout)
	return nil
}

func (source *userAccountToken) refresh(ctx context.Context) error {
	resp, err := source.CommonClient.RefreshJwt(ctx, &jwt.RefreshJwtRequest{
		Refresh: source.refreshToken,
	})
	if err != nil {
		// refresh失败，尝试重新登录
		err := source.init(ctx)
		if err != nil {
			return err
		}
		return err
	}
	source.mutex.Lock()
	defer source.mutex.Unlock()
	source.accessToken = resp.Access
	source.refreshToken = resp.Refresh
	source.expiry = time.Now().Add(source.opts.tokenTimeout)
	return nil
}

func (source *userAccountToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := source.tokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"Authorization": token.Type() + " " + token.AccessToken,
	}, nil
}

// 直接访问furnace grpc接口，无需安全证书
func (source *userAccountToken) RequireTransportSecurity() bool {
	return false
}
