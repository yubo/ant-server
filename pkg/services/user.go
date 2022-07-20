package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yubo/ant-server/pkg/api"
	"github.com/yubo/ant-server/pkg/models"
	"github.com/yubo/apiserver/pkg/options"
	"github.com/yubo/apiserver/pkg/rest"
	"github.com/yubo/golib/util"
)

type user struct {
	user *models.User
}

func newUser() *user {
	return &user{
		user: models.NewUser(),
	}
}

// Because some configuration may be stored in the database,
func (p *user) install(ctx context.Context) error {
	http, ok := options.APIServerFrom(ctx)
	if !ok {
		return fmt.Errorf("unable to get API server from the context")
	}

	p.installWs(http)

	return nil
}

func (p *user) installWs(container rest.GoRestfulContainer) {
	rest.SwaggerTagRegister("user", "user Api")
	rest.WsRouteBuild(&rest.WsOption{
		Path:               "/api/v1/users",
		Tags:               []string{"user"},
		GoRestfulContainer: container,
		Routes: []rest.WsRoute{{
			Method: "POST", SubPath: "/",
			Desc:      "create user",
			Operation: "createUser",
			Handle:    p.create,
		}, {
			Method: "GET", SubPath: "/",
			Desc:      "search/list users",
			Operation: "listUser",
			Handle:    p.listUser,
		}, {
			Method: "GET", SubPath: "/{name}",
			Desc:      "get user",
			Operation: "getUser",
			Handle:    p.getUser,
		}, {
			Method: "PUT", SubPath: "/",
			Desc:      "update user",
			Operation: "updateUser",
			Handle:    p.updateUser,
		}, {
			Method: "DELETE", SubPath: "/{name}",
			Desc:      "delete user",
			Operation: "deleteUser",
			Handle:    p.deleteUser,
		}},
	})
}

// User CRUD
func (p *user) create(w http.ResponseWriter, req *http.Request, _ *rest.NonParam, in *api.CreateUserInput) (*api.User, error) {
	if err := p.user.Create(req.Context(), in); err != nil {
		return nil, err
	}

	return p.user.Get(req.Context(), util.StringValue(in.Name))
}

func (p *user) getUser(w http.ResponseWriter, req *http.Request, in *api.NameParam) (*api.User, error) {
	return p.user.Get(req.Context(), in.Name)
}

func (p *user) listUser(w http.ResponseWriter, req *http.Request, in *api.ListParam) (ret *api.ListUserOutput, err error) {
	ret = &api.ListUserOutput{}

	opts, err := in.ListOptions(in.Query, &ret.Total)
	if err != nil {
		return nil, err
	}

	ret.List, err = p.user.List(req.Context(), *opts)
	return ret, err
}

func (p *user) updateUser(w http.ResponseWriter, req *http.Request, param *rest.NonParam, in *api.UpdateUserInput) error {
	return p.user.Update(req.Context(), in)
}

func (p *user) deleteUser(w http.ResponseWriter, req *http.Request, in *api.NameParam) error {
	return p.user.Delete(req.Context(), in.Name)
}
