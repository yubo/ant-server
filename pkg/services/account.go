package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yubo/ant-server/pkg/api"
	"github.com/yubo/apiserver/pkg/options"
	"github.com/yubo/apiserver/pkg/request"
	"github.com/yubo/apiserver/pkg/rest"
	"github.com/yubo/apiserver/pkg/server"
	"github.com/yubo/golib/proc"

	_ "github.com/yubo/apiserver/pkg/models/register"
	_ "github.com/yubo/apiserver/pkg/session/register"
)

func init() {
	proc.AddConfig("account", newAccountConfig())
}

// {{{ accountConfig
type accountConfig struct {
}

func newAccountConfig() *accountConfig {
	return &accountConfig{}
}

// }}}

type account struct {
	config *accountConfig
}

func newAccount() *account {
	return &account{}
}

// Because some configuration may be stored in the database,
func (p *account) install(ctx context.Context) error {
	config := newAccountConfig()
	if err := proc.ReadConfig("account", config); err != nil {
		return err
	}
	p.config = config

	http, ok := options.APIServerFrom(ctx)
	if !ok {
		return fmt.Errorf("unable to get API server from the context")
	}
	p.installWs(http)

	return nil
}

func (p *account) installWs(container server.APIServer) {
	rest.SwaggerTagRegister("account", "account Api")
	rest.WsRouteBuild(&rest.WsOption{
		Path:               "/api/v1/account",
		Tags:               []string{"account"},
		GoRestfulContainer: container,
		Routes: []rest.WsRoute{{
			Desc:   "account login",
			Method: "POST", SubPath: "/login",
			Operation: "accountLogin",
			Handle:    p.login,
		}, {
			Desc:   "发送验证码",
			Method: "POST", SubPath: "/captcha",
			Operation: "accountCaptcha",
			Handle:    p.captcha,
		}, {
			Desc:   "get curent user info",
			Method: "GET", SubPath: "/info",
			Operation: "accountInfo",
			Handle:    p.getCurrentInfo,
		}, {
			Desc:   "account logout",
			Method: "POST", SubPath: "/logout",
			Operation: "accountLogout",
			Handle:    p.logout,
		}},
	})
}

func (p *account) login(w http.ResponseWriter, req *http.Request, _ *rest.NonParam, in *api.LoginInput) (*api.LoginOutput, error) {
	sess, ok := request.SessionFrom(req.Context())
	if !ok {
		return nil, fmt.Errorf("cat't get session")
	}

	sess.Set("userName", "tom")
	sess.Set("groups", "dev,admin,ops")

	return &api.LoginOutput{
		Type:             in.Type,
		CurrentAuthority: "rbac",
	}, nil
}

func (p *account) captcha(w http.ResponseWriter, req *http.Request, _ *rest.NonParam, in *api.CaptchaInput) error {
	return nil
}

func (p *account) getCurrentInfo(w http.ResponseWriter, req *http.Request) (*api.User, error) {
	return &api.User{
		Name:  "tom",
		Group: "dev,admin,ops",
	}, nil
}

func (p *account) logout(w http.ResponseWriter, req *http.Request) error {
	sess, ok := request.SessionFrom(req.Context())
	if ok {
		sess.Reset()
	}
	return nil
}
